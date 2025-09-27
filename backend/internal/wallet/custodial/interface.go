package custodial

import "context"

// Wallet represents a custodial wallet
type Wallet struct {
    ID       string
    Address  string
    Chain    string
    Balance  float64
    IsActive bool
}

// Transaction represents a custodial wallet transaction
type Transaction struct {
    ID            string
    WalletID      string
    TxHash        string
    FromAddress   string
    ToAddress     string
    Amount        float64
    Chain         string
    Status        string
    Confirmations int
    Fee           float64
    CreatedAt     int64
}

// Provider defines the interface for custodial wallet providers
type Provider interface {
    // CreateWallet creates a new custodial wallet
    CreateWallet(ctx context.Context, chain string) (*Wallet, error)
    
    // GetWallet retrieves wallet information
    GetWallet(ctx context.Context, walletID string) (*Wallet, error)
    
    // GetWalletByAddress retrieves wallet information by address
    GetWalletByAddress(ctx context.Context, address string) (*Wallet, error)
    
    // GetBalance retrieves the balance of a wallet
    GetBalance(ctx context.Context, walletID string) (float64, error)
    
    // SendTransaction sends a transaction from a custodial wallet
    SendTransaction(ctx context.Context, walletID, to string, amount float64) (*Transaction, error)
    
    // GetTransaction retrieves transaction details
    GetTransaction(ctx context.Context, txID string) (*Transaction, error)
    
    // ListTransactions lists transactions for a wallet
    ListTransactions(ctx context.Context, walletID string, limit, offset int) ([]Transaction, error)
    
    // GetSupportedChains returns the list of supported blockchain networks
    GetSupportedChains() []string
}