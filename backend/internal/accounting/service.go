package accounting

import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// LedgerService handles accounting and ledger operations
type LedgerService struct {
    db *gorm.DB
}

// NewLedgerService creates a new ledger service
func NewLedgerService(db *gorm.DB) *LedgerService {
    return &LedgerService{
        db: db,
    }
}

// InitializeDB initializes the database tables for accounting service
func (s *LedgerService) InitializeDB() error {
    return s.db.AutoMigrate(&Account{}, &Transaction{}, &JournalEntry{}, &Balance{})
}

// CreateAccount creates a new account
func (s *LedgerService) CreateAccount(ctx context.Context, account Account) (*Account, error) {
    account.CreatedAt = time.Now()
    account.UpdatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(&account).Error; err != nil {
        return nil, fmt.Errorf("failed to create account: %w", err)
    }
    
    return &account, nil
}

// GetAccount retrieves an account by ID
func (s *LedgerService) GetAccount(ctx context.Context, accountID uint) (*Account, error) {
    var account Account
    if err := s.db.WithContext(ctx).First(&account, "id = ?", accountID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("account not found")
        }
        return nil, fmt.Errorf("failed to get account: %w", err)
    }
    
    return &account, nil
}

// GetAccountByUserID retrieves an account by user ID
func (s *LedgerService) GetAccountByUserID(ctx context.Context, userID uint) (*Account, error) {
    var account Account
    if err := s.db.WithContext(ctx).First(&account, "user_id = ?", userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("account not found")
        }
        return nil, fmt.Errorf("failed to get account: %w", err)
    }
    
    return &account, nil
}

// UpdateAccountBalance updates an account's balance
func (s *LedgerService) UpdateAccountBalance(ctx context.Context, accountID uint, amount float64) error {
    return s.db.WithContext(ctx).Model(&Account{}).Where("id = ?", accountID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

// FreezeAccount freezes an account
func (s *LedgerService) FreezeAccount(ctx context.Context, accountID uint) error {
    return s.db.WithContext(ctx).Model(&Account{}).Where("id = ?", accountID).Update("is_frozen", true).Error
}

// UnfreezeAccount unfreezes an account
func (s *LedgerService) UnfreezeAccount(ctx context.Context, accountID uint) error {
    return s.db.WithContext(ctx).Model(&Account{}).Where("id = ?", accountID).Update("is_frozen", false).Error
}

// CreateTransaction creates a new transaction
func (s *LedgerService) CreateTransaction(ctx context.Context, transaction Transaction) (*Transaction, error) {
    transaction.CreatedAt = time.Now()
    transaction.UpdatedAt = time.Now()
    transaction.Status = StatusPending
    
    if err := s.db.WithContext(ctx).Create(&transaction).Error; err != nil {
        return nil, fmt.Errorf("failed to create transaction: %w", err)
    }
    
    return &transaction, nil
}

// GetTransaction retrieves a transaction by ID
func (s *LedgerService) GetTransaction(ctx context.Context, transactionID uint) (*Transaction, error) {
    var transaction Transaction
    if err := s.db.WithContext(ctx).First(&transaction, "id = ?", transactionID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("transaction not found")
        }
        return nil, fmt.Errorf("failed to get transaction: %w", err)
    }
    
    return &transaction, nil
}

// GetTransactionByReference retrieves a transaction by reference ID
func (s *LedgerService) GetTransactionByReference(ctx context.Context, referenceID string) (*Transaction, error) {
    var transaction Transaction
    if err := s.db.WithContext(ctx).First(&transaction, "reference_id = ?", referenceID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("transaction not found")
        }
        return nil, fmt.Errorf("failed to get transaction: %w", err)
    }
    
    return &transaction, nil
}

// UpdateTransactionStatus updates a transaction's status
func (s *LedgerService) UpdateTransactionStatus(ctx context.Context, transactionID uint, status TransactionStatus) error {
    return s.db.WithContext(ctx).Model(&Transaction{}).Where("id = ?", transactionID).Update("status", status).Error
}

// CreateJournalEntry creates a new journal entry
func (s *LedgerService) CreateJournalEntry(ctx context.Context, entry JournalEntry) (*JournalEntry, error) {
    entry.CreatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(&entry).Error; err != nil {
        return nil, fmt.Errorf("failed to create journal entry: %w", err)
    }
    
    return &entry, nil
}

// GetJournalEntriesByTransaction retrieves journal entries for a transaction
func (s *LedgerService) GetJournalEntriesByTransaction(ctx context.Context, transactionID uint) ([]JournalEntry, error) {
    var entries []JournalEntry
    if err := s.db.WithContext(ctx).Where("transaction_id = ?", transactionID).Find(&entries).Error; err != nil {
        return nil, fmt.Errorf("failed to get journal entries: %w", err)
    }
    
    return entries, nil
}

// GetAccountBalance retrieves the current balance for an account
func (s *LedgerService) GetAccountBalance(ctx context.Context, accountID uint) (float64, error) {
    var account Account
    if err := s.db.WithContext(ctx).First(&account, "id = ?", accountID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return 0, fmt.Errorf("account not found")
        }
        return 0, fmt.Errorf("failed to get account: %w", err)
    }
    
    return account.Balance, nil
}

// GetAccountTransactions retrieves transactions for an account
func (s *LedgerService) GetAccountTransactions(ctx context.Context, accountID uint, limit, offset int) ([]Transaction, error) {
    var transactions []Transaction
    if err := s.db.WithContext(ctx).Where("account_id = ?", accountID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&transactions).Error; err != nil {
        return nil, fmt.Errorf("failed to get transactions: %w", err)
    }
    
    return transactions, nil
}