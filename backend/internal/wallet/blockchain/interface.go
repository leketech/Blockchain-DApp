package blockchain

import "context"

// Wallet represents a blockchain wallet
type Wallet struct {
    Address    string
    PublicKey  string
    PrivateKey string
    Balance    float64
}

// Transaction represents a blockchain transaction
type Transaction struct {
    Hash          string
    From          string
    To            string
    Amount        float64
    Fee           float64
    Confirmations int
    Status        string // pending, confirmed, failed
    Timestamp     int64
}

// Adapter defines the interface for blockchain adapters
type Adapter interface {
    // CreateWallet creates a new wallet
    CreateWallet(ctx context.Context) (*Wallet, error)
    
    // GetWallet retrieves wallet information
    GetWallet(ctx context.Context, address string) (*Wallet, error)
    
    // GetBalance retrieves the balance of an address
    GetBalance(ctx context.Context, address string) (float64, error)
    
    // SendTransaction sends a transaction
    SendTransaction(ctx context.Context, from, to string, amount float64, privateKey string) (*Transaction, error)
    
    // GetTransaction retrieves transaction details
    GetTransaction(ctx context.Context, hash string) (*Transaction, error)
    
    // EstimateFee estimates the transaction fee
    EstimateFee(ctx context.Context, from, to string, amount float64) (float64, error)
}