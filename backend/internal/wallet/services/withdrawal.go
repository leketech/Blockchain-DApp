package services

import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/blockchain-dapp/backend/internal/wallet"
    "github.com/blockchain-dapp/backend/internal/wallet/blockchain"
    "github.com/blockchain-dapp/backend/internal/wallet/custodial"
    "gorm.io/gorm"
)

// WithdrawalService handles withdrawal requests using custodial wallets or direct signing
type WithdrawalService struct {
    db              *gorm.DB
    adapters        map[string]blockchain.Adapter
    custodialProviders map[string]custodial.Provider
    mu              sync.RWMutex
}

// NewWithdrawalService creates a new withdrawal service
func NewWithdrawalService(db *gorm.DB, adapters map[string]blockchain.Adapter, custodialProviders map[string]custodial.Provider) *WithdrawalService {
    return &WithdrawalService{
        db:                 db,
        adapters:           adapters,
        custodialProviders: custodialProviders,
        mu:                 sync.RWMutex{},
    }
}

// RequestWithdrawal creates a new withdrawal request
func (ws *WithdrawalService) RequestWithdrawal(ctx context.Context, userID uint, chain, toAddress string, amount float64, useCustodial bool) (*wallet.Transaction, error) {
    // Validate the destination address
    if !ws.isValidAddress(chain, toAddress) {
        return nil, fmt.Errorf("invalid destination address for chain %s", chain)
    }
    
    // Create a transaction record
    transaction := &wallet.Transaction{
        UserID:      userID,
        Chain:       chain,
        ToAddress:   toAddress,
        Amount:      amount,
        Status:      "pending",
        CreatedAt:   time.Now(),
        UseCustodial: useCustodial,
    }
    
    // Save the transaction to the database
    if err := ws.db.Create(transaction).Error; err != nil {
        return nil, fmt.Errorf("failed to save transaction: %w", err)
    }
    
    return transaction, nil
}

// ProcessWithdrawal processes a withdrawal request
func (ws *WithdrawalService) ProcessWithdrawal(ctx context.Context, transactionID uint) error {
    // Get the transaction
    var transaction wallet.Transaction
    err := ws.db.First(&transaction, transactionID).Error
    if err != nil {
        return fmt.Errorf("failed to fetch transaction: %w", err)
    }
    
    // Check if the transaction is already processed
    if transaction.Status != "pending" {
        return fmt.Errorf("transaction is not in pending status")
    }
    
    // Process based on whether we're using custodial or direct signing
    if transaction.UseCustodial {
        return ws.processCustodialWithdrawal(ctx, &transaction)
    }
    
    return ws.processDirectWithdrawal(ctx, &transaction)
}

// processCustodialWithdrawal processes a withdrawal using a custodial provider
func (ws *WithdrawalService) processCustodialWithdrawal(ctx context.Context, transaction *wallet.Transaction) error {
    ws.mu.RLock()
    provider, exists := ws.custodialProviders[transaction.Chain]
    ws.mu.RUnlock()
    
    if !exists {
        return fmt.Errorf("no custodial provider available for chain %s", transaction.Chain)
    }
    
    // Get the user's custodial wallet
    custodialWallet, err := ws.getCustodialWallet(ctx, transaction.UserID, transaction.Chain)
    if err != nil {
        return fmt.Errorf("failed to get custodial wallet: %w", err)
    }
    
    // Send the transaction through the custodial provider
    tx, err := provider.SendTransaction(ctx, custodialWallet.ExternalID, transaction.ToAddress, transaction.Amount)
    if err != nil {
        transaction.Status = "failed"
        transaction.ErrorMessage = err.Error()
        ws.db.Save(transaction)
        return fmt.Errorf("failed to send transaction through custodial provider: %w", err)
    }
    
    // Update the transaction record
    transaction.TxHash = tx.TxHash
    transaction.FromAddress = tx.FromAddress
    transaction.Fee = tx.Fee
    transaction.Status = tx.Status
    transaction.Confirmations = tx.Confirmations
    transaction.UpdatedAt = time.Now()
    
    return ws.db.Save(transaction).Error
}

// processDirectWithdrawal processes a withdrawal using direct signing
func (ws *WithdrawalService) processDirectWithdrawal(ctx context.Context, transaction *wallet.Transaction) error {
    // Security review: This is a critical function that handles private key operations
    // In a production environment, this should be handled with extreme care and proper
    // security measures such as HSMs, secure key storage, and multi-signature schemes
    
    // SECURITY REVIEW:
    // 1. Private keys should never be stored in plain text in the database
    // 2. All signing operations should be performed in a secure environment (HSM)
    // 3. Multi-signature schemes should be used for high-value transactions
    // 4. Rate limiting should be implemented to prevent abuse
    // 5. All transactions should be logged for audit purposes
    // 6. Two-factor authentication should be required for withdrawals
    // 7. Withdrawal limits should be enforced per user and per time period
    
    ws.mu.RLock()
    adapter, exists := ws.adapters[transaction.Chain]
    ws.mu.RUnlock()
    
    if !exists {
        return fmt.Errorf("unsupported chain: %s", transaction.Chain)
    }
    
    // Get the user's private wallet
    w, err := ws.getPrivateWallet(ctx, transaction.UserID, transaction.Chain)
    if err != nil {
        return fmt.Errorf("failed to get private wallet: %w", err)
    }
    
    // Send the transaction using the blockchain adapter
    // Note: In a real implementation, the private key should never be stored in plain text
    // and should be retrieved from a secure key management system
    tx, err := adapter.SendTransaction(ctx, w.Address, transaction.ToAddress, transaction.Amount, w.PrivateKey)
    if err != nil {
        transaction.Status = "failed"
        transaction.ErrorMessage = err.Error()
        ws.db.Save(transaction)
        return fmt.Errorf("failed to send transaction: %w", err)
    }
    
    // Update the transaction record
    transaction.TxHash = tx.Hash
    transaction.FromAddress = tx.From
    transaction.Fee = tx.Fee
    transaction.Status = tx.Status
    transaction.Confirmations = tx.Confirmations
    transaction.UpdatedAt = time.Now()
    
    return ws.db.Save(transaction).Error
}

// getCustodialWallet retrieves a user's custodial wallet for a specific chain
func (ws *WithdrawalService) getCustodialWallet(ctx context.Context, userID uint, chain string) (*wallet.CustodialWallet, error) {
    var w wallet.CustodialWallet
    err := ws.db.Where("user_id = ? AND chain = ?", userID, chain).First(&w).Error
    if err != nil {
        return nil, err
    }
    
    return &w, nil
}

// getPrivateWallet retrieves a user's private wallet for a specific chain
func (ws *WithdrawalService) getPrivateWallet(ctx context.Context, userID uint, chain string) (*wallet.Wallet, error) {
    var w wallet.Wallet
    err := ws.db.Where("user_id = ? AND chain = ? AND type = ?", userID, chain, "private").First(&w).Error
    if err != nil {
        return nil, err
    }
    
    return &w, nil
}

// isValidAddress validates an address for a specific chain
func (ws *WithdrawalService) isValidAddress(chain, address string) bool {
    ws.mu.RLock()
    adapter, exists := ws.adapters[chain]
    ws.mu.RUnlock()
    
    if !exists {
        return false
    }
    
    // In a real implementation, each adapter would have its own address validation logic
    // For now, we'll just check that the address is not empty
    return address != ""
}

// CancelWithdrawal cancels a pending withdrawal request
func (ws *WithdrawalService) CancelWithdrawal(ctx context.Context, transactionID uint) error {
    var transaction wallet.Transaction
    err := ws.db.First(&transaction, transactionID).Error
    if err != nil {
        return fmt.Errorf("failed to fetch transaction: %w", err)
    }
    
    if transaction.Status != "pending" {
        return fmt.Errorf("transaction is not in pending status")
    }
    
    transaction.Status = "cancelled"
    transaction.UpdatedAt = time.Now()
    
    return ws.db.Save(&transaction).Error
}

// SecurityReview returns the security review checklist for the withdrawal service
func (ws *WithdrawalService) SecurityReview() map[string]bool {
    return map[string]bool{
        "Private keys encrypted at rest": false, // Should be true in production
        "Signing operations in secure environment (HSM)": false, // Should be true in production
        "Multi-signature for high-value transactions": false, // Should be implemented
        "Rate limiting for withdrawal requests": false, // Should be implemented
        "Audit logging for all transactions": true, // Implemented through database records
        "Two-factor authentication for withdrawals": false, // Should be integrated with auth service
        "Withdrawal limits enforced": false, // Should be implemented
        "Custodial provider integration secure": true, // Using established providers
    }
}