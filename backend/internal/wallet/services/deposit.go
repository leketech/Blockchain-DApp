package services

import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/blockchain-dapp/backend/internal/wallet"
    "github.com/blockchain-dapp/backend/internal/wallet/blockchain"
    "gorm.io/gorm"
)

// DepositService handles deposit address generation and monitoring
type DepositService struct {
    db       *gorm.DB
    adapters map[string]blockchain.Adapter
    mu       sync.RWMutex
}

// NewDepositService creates a new deposit service
func NewDepositService(db *gorm.DB, adapters map[string]blockchain.Adapter) *DepositService {
    return &DepositService{
        db:       db,
        adapters: adapters,
        mu:       sync.RWMutex{},
    }
}

// GenerateDepositAddress generates a new deposit address for a user and chain
func (s *DepositService) GenerateDepositAddress(ctx context.Context, userID uint, chain string) (*wallet.Wallet, error) {
    s.mu.RLock()
    adapter, exists := s.adapters[chain]
    s.mu.RUnlock()
    
    if !exists {
        return nil, fmt.Errorf("unsupported chain: %s", chain)
    }
    
    // Create a new wallet using the blockchain adapter
    w, err := adapter.CreateWallet(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create wallet: %w", err)
    }
    
    // Save the wallet to the database
    depositWallet := &wallet.Wallet{
        UserID:    userID,
        Chain:     chain,
        Address:   w.Address,
        PublicKey: w.PublicKey,
        Balance:   w.Balance,
        Type:      "deposit",
    }
    
    if err := s.db.Create(depositWallet).Error; err != nil {
        return nil, fmt.Errorf("failed to save wallet to database: %w", err)
    }
    
    return depositWallet, nil
}

// GetDepositAddress retrieves a user's deposit address for a specific chain
func (s *DepositService) GetDepositAddress(ctx context.Context, userID uint, chain string) (*wallet.Wallet, error) {
    var w wallet.Wallet
    err := s.db.Where("user_id = ? AND chain = ? AND type = ?", userID, chain, "deposit").First(&w).Error
    if err != nil {
        return nil, err
    }
    
    return &w, nil
}

// GetAllDepositAddresses retrieves all deposit addresses for a user
func (s *DepositService) GetAllDepositAddresses(ctx context.Context, userID uint) ([]wallet.Wallet, error) {
    var wallets []wallet.Wallet
    err := s.db.Where("user_id = ? AND type = ?", userID, "deposit").Find(&wallets).Error
    if err != nil {
        return nil, err
    }
    
    return wallets, nil
}

// UpdateDepositAddressBalance updates the balance of a deposit address
func (s *DepositService) UpdateDepositAddressBalance(ctx context.Context, walletID uint) error {
    var w wallet.Wallet
    err := s.db.First(&w, walletID).Error
    if err != nil {
        return err
    }
    
    s.mu.RLock()
    adapter, exists := s.adapters[w.Chain]
    s.mu.RUnlock()
    
    if !exists {
        return fmt.Errorf("unsupported chain: %s", w.Chain)
    }
    
    // Get the current balance from the blockchain
    balance, err := adapter.GetBalance(ctx, w.Address)
    if err != nil {
        return fmt.Errorf("failed to get balance: %w", err)
    }
    
    // Update the wallet in the database
    w.Balance = balance
    w.UpdatedAt = time.Now()
    
    return s.db.Save(&w).Error
}