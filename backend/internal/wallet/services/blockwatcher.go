package services

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "sync"
    "time"

    "github.com/blockchain-dapp/backend/internal/wallet"
    "github.com/blockchain-dapp/backend/internal/wallet/blockchain"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "gorm.io/gorm"
)

// BlockWatcher monitors blockchain blocks for deposit transactions
type BlockWatcher struct {
    db           *gorm.DB
    client       *ethclient.Client
    chain        string
    lastBlock    *big.Int
    pollInterval time.Duration
    mu           sync.RWMutex
    running      bool
}

// NewBlockWatcher creates a new block watcher
func NewBlockWatcher(db *gorm.DB, rpcURL, chain string, pollInterval time.Duration) (*BlockWatcher, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
    }
    
    return &BlockWatcher{
        db:           db,
        client:       client,
        chain:        chain,
        pollInterval: pollInterval,
        mu:           sync.RWMutex{},
    }, nil
}

// Start begins watching for new blocks
func (bw *BlockWatcher) Start(ctx context.Context) error {
    bw.mu.Lock()
    if bw.running {
        bw.mu.Unlock()
        return fmt.Errorf("block watcher is already running")
    }
    bw.running = true
    bw.mu.Unlock()
    
    // Get the current block number to start watching from
    header, err := bw.client.HeaderByNumber(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to get current block header: %w", err)
    }
    
    bw.lastBlock = header.Number
    
    log.Printf("Starting block watcher for %s from block %s", bw.chain, bw.lastBlock.String())
    
    // Start the watching loop
    ticker := time.NewTicker(bw.pollInterval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            bw.mu.Lock()
            bw.running = false
            bw.mu.Unlock()
            return ctx.Err()
        case <-ticker.C:
            if err := bw.watchBlocks(ctx); err != nil {
                log.Printf("Error watching blocks: %v", err)
            }
        }
    }
}

// watchBlocks checks for new blocks and processes transactions
func (bw *BlockWatcher) watchBlocks(ctx context.Context) error {
    // Get the current block number
    header, err := bw.client.HeaderByNumber(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to get current block header: %w", err)
    }
    
    currentBlock := header.Number
    
    // If we haven't processed any blocks yet, start from the current block
    if bw.lastBlock == nil {
        bw.lastBlock = currentBlock
        return nil
    }
    
    // Process blocks from lastBlock + 1 to currentBlock
    for i := new(big.Int).Add(bw.lastBlock, big.NewInt(1)); i.Cmp(currentBlock) <= 0; i.Add(i, big.NewInt(1)) {
        if err := bw.processBlock(ctx, i); err != nil {
            log.Printf("Error processing block %s: %v", i.String(), err)
        }
    }
    
    // Update the last processed block
    bw.lastBlock = currentBlock
    
    return nil
}

// processBlock processes transactions in a block
func (bw *BlockWatcher) processBlock(ctx context.Context, blockNumber *big.Int) error {
    log.Printf("Processing block %s", blockNumber.String())
    
    // Get the block by number
    block, err := bw.client.BlockByNumber(ctx, blockNumber)
    if err != nil {
        return fmt.Errorf("failed to get block %s: %w", blockNumber.String(), err)
    }
    
    // Process each transaction in the block
    for _, tx := range block.Transactions() {
        if err := bw.processTransaction(ctx, tx); err != nil {
            log.Printf("Error processing transaction %s: %v", tx.Hash().Hex(), err)
        }
    }
    
    return nil
}

// processTransaction checks if a transaction is a deposit to one of our addresses
func (bw *BlockWatcher) processTransaction(ctx context.Context, tx *types.Transaction) error {
    // We only care about transactions that have a recipient (not contract creation)
    if tx.To() == nil {
        return nil
    }
    
    // Check if the recipient address is one of our deposit addresses
    var wallet wallet.Wallet
    err := bw.db.Where("address = ? AND chain = ?", tx.To().Hex(), bw.chain).First(&wallet).Error
    if err != nil {
        // Not one of our addresses, that's fine
        if err == gorm.ErrRecordNotFound {
            return nil
        }
        return fmt.Errorf("failed to query wallet: %w", err)
    }
    
    // This is a deposit to one of our addresses
    log.Printf("Found deposit transaction %s to address %s", tx.Hash().Hex(), tx.To().Hex())
    
    // Get the transaction receipt to confirm it
    receipt, err := bw.client.TransactionReceipt(ctx, tx.Hash())
    if err != nil {
        return fmt.Errorf("failed to get transaction receipt: %w", err)
    }
    
    // Check if the transaction was successful
    if receipt.Status != types.ReceiptStatusSuccessful {
        log.Printf("Transaction %s failed, not processing", tx.Hash().Hex())
        return nil
    }
    
    // Create a transaction record
    transaction := &wallet.Transaction{
        WalletID:    wallet.ID,
        TxHash:      tx.Hash().Hex(),
        FromAddress: "", // Would need to derive from signature
        ToAddress:   tx.To().Hex(),
        Amount:      float64(tx.Value().Int64()) / 1000000000000000000, // Convert wei to ETH
        Chain:       bw.chain,
        Status:      "confirmed",
        Confirmations: int(receipt.BlockNumber.Uint64()),
        Fee:         0, // Would calculate from gas price and limit
    }
    
    // Save the transaction to the database
    if err := bw.db.Create(transaction).Error; err != nil {
        return fmt.Errorf("failed to save transaction: %w", err)
    }
    
    // Update the wallet balance
    wallet.Balance += transaction.Amount
    if err := bw.db.Save(&wallet).Error; err != nil {
        return fmt.Errorf("failed to update wallet balance: %w", err)
    }
    
    log.Printf("Processed deposit transaction %s for %f ETH to wallet %d", tx.Hash().Hex(), transaction.Amount, wallet.ID)
    
    return nil
}

// Stop stops the block watcher
func (bw *BlockWatcher) Stop() {
    bw.mu.Lock()
    defer bw.mu.Unlock()
    bw.running = false
}

// IsRunning returns whether the block watcher is currently running
func (bw *BlockWatcher) IsRunning() bool {
    bw.mu.RLock()
    defer bw.mu.RUnlock()
    return bw.running
}

// ConnectToTestnet connects the block watcher to an Ethereum testnet
func (bw *BlockWatcher) ConnectToTestnet(network string) error {
    var rpcURL string
    switch network {
    case "goerli":
        rpcURL = "https://goerli.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    case "sepolia":
        rpcURL = "https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    case "rinkeby":
        rpcURL = "https://rinkeby.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    default:
        return fmt.Errorf("unsupported testnet: %s", network)
    }
    
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return fmt.Errorf("failed to connect to %s testnet: %w", network, err)
    }
    
    bw.client = client
    
    return nil
}