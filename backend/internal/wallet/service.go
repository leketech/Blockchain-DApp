package wallet

import (
    "errors"

    "gorm.io/gorm"
)

// Service provides wallet operations
type Service struct {
    db *gorm.DB
}

// NewService creates a new wallet service
func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

// CreateWallet creates a new wallet
func (s *Service) CreateWallet(wallet *Wallet) error {
    if wallet.Address == "" || wallet.Chain == "" || wallet.PublicKey == "" {
        return errors.New("address, chain, and public key are required")
    }

    return s.db.Create(wallet).Error
}

// GetWalletByID retrieves a wallet by ID
func (s *Service) GetWalletByID(id uint) (*Wallet, error) {
    var wallet Wallet
    err := s.db.First(&wallet, id).Error
    if err != nil {
        return nil, err
    }
    return &wallet, nil
}

// GetWalletByAddress retrieves a wallet by address
func (s *Service) GetWalletByAddress(address string) (*Wallet, error) {
    var wallet Wallet
    err := s.db.Where("address = ?", address).First(&wallet).Error
    if err != nil {
        return nil, err
    }
    return &wallet, nil
}

// GetWalletsByUserID retrieves all wallets for a user
func (s *Service) GetWalletsByUserID(userID uint) ([]Wallet, error) {
    var wallets []Wallet
    err := s.db.Where("user_id = ?", userID).Find(&wallets).Error
    return wallets, err
}

// UpdateWallet updates a wallet
func (s *Service) UpdateWallet(wallet *Wallet) error {
    return s.db.Save(wallet).Error
}

// DeleteWallet deletes a wallet
func (s *Service) DeleteWallet(id uint) error {
    return s.db.Delete(&Wallet{}, id).Error
}

// CreateTransaction creates a new transaction
func (s *Service) CreateTransaction(tx *Transaction) error {
    if tx.TxHash == "" || tx.FromAddress == "" || tx.ToAddress == "" || tx.Amount <= 0 {
        return errors.New("transaction hash, from address, to address, and amount are required")
    }

    return s.db.Create(tx).Error
}

// GetTransactionByHash retrieves a transaction by hash
func (s *Service) GetTransactionByHash(hash string) (*Transaction, error) {
    var tx Transaction
    err := s.db.Where("tx_hash = ?", hash).First(&tx).Error
    if err != nil {
        return nil, err
    }
    return &tx, nil
}

// GetTransactionsByWalletID retrieves all transactions for a wallet
func (s *Service) GetTransactionsByWalletID(walletID uint) ([]Transaction, error) {
    var transactions []Transaction
    err := s.db.Where("wallet_id = ?", walletID).Find(&transactions).Error
    return transactions, err
}

// UpdateTransaction updates a transaction
func (s *Service) UpdateTransaction(tx *Transaction) error {
    return s.db.Save(tx).Error
}

// CreateCustodialWallet creates a new custodial wallet
func (s *Service) CreateCustodialWallet(wallet *CustodialWallet) error {
    if wallet.ExternalID == "" || wallet.Provider == "" || wallet.Chain == "" || wallet.Address == "" {
        return errors.New("external ID, provider, chain, and address are required")
    }

    return s.db.Create(wallet).Error
}

// GetCustodialWalletByID retrieves a custodial wallet by ID
func (s *Service) GetCustodialWalletByID(id uint) (*CustodialWallet, error) {
    var wallet CustodialWallet
    err := s.db.First(&wallet, id).Error
    if err != nil {
        return nil, err
    }
    return &wallet, nil
}

// GetCustodialWalletByExternalID retrieves a custodial wallet by external ID
func (s *Service) GetCustodialWalletByExternalID(externalID string) (*CustodialWallet, error) {
    var wallet CustodialWallet
    err := s.db.Where("external_id = ?", externalID).First(&wallet).Error
    if err != nil {
        return nil, err
    }
    return &wallet, nil
}

// GetCustodialWalletsByUserID retrieves all custodial wallets for a user
func (s *Service) GetCustodialWalletsByUserID(userID uint) ([]CustodialWallet, error) {
    var wallets []CustodialWallet
    err := s.db.Where("user_id = ?", userID).Find(&wallets).Error
    return wallets, err
}