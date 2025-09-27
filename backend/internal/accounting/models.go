package accounting

import (
    "time"
)

// TransactionType represents the type of financial transaction
type TransactionType string

const (
    TypeDeposit     TransactionType = "deposit"
    TypeWithdrawal  TransactionType = "withdrawal"
    TypeTransfer    TransactionType = "transfer"
    TypeFee         TransactionType = "fee"
    TypeInterest    TransactionType = "interest"
    TypeAdjustment  TransactionType = "adjustment"
    TypePayment     TransactionType = "payment"
    TypeRefund      TransactionType = "refund"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
    StatusPending   TransactionStatus = "pending"
    StatusCompleted TransactionStatus = "completed"
    StatusFailed    TransactionStatus = "failed"
    StatusCancelled TransactionStatus = "cancelled"
    StatusReversed  TransactionStatus = "reversed"
)

// AccountType represents the type of account
type AccountType string

const (
    AccountTypeAsset     AccountType = "asset"
    AccountTypeLiability AccountType = "liability"
    AccountTypeEquity    AccountType = "equity"
    AccountTypeRevenue   AccountType = "revenue"
    AccountTypeExpense   AccountType = "expense"
)

// Account represents a financial account
type Account struct {
    ID          uint        `gorm:"primaryKey" json:"id"`
    UserID      *uint       `json:"user_id,omitempty"`
    Name        string      `gorm:"not null" json:"name"`
    Type        AccountType `gorm:"not null" json:"type"`
    Currency    string      `gorm:"not null" json:"currency"`
    Balance     float64     `gorm:"not null;default:0" json:"balance"`
    IsFrozen    bool        `gorm:"not null;default:false" json:"is_frozen"`
    Description string      `json:"description,omitempty"`
    Metadata    string      `json:"metadata,omitempty"` // JSON field for additional data
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

// Transaction represents a financial transaction
type Transaction struct {
    ID              uint             `gorm:"primaryKey" json:"id"`
    ReferenceID     string           `gorm:"uniqueIndex;not null" json:"reference_id"`
    AccountID       uint             `gorm:"not null" json:"account_id"`
    CounterpartyID  *uint            `json:"counterparty_id,omitempty"`
    Type            TransactionType  `gorm:"not null" json:"type"`
    Status          TransactionStatus `gorm:"not null" json:"status"`
    Amount          float64          `gorm:"not null" json:"amount"`
    Currency        string           `gorm:"not null" json:"currency"`
    Fee             *float64         `json:"fee,omitempty"`
    Description     string           `json:"description,omitempty"`
    Metadata        string           `json:"metadata,omitempty"` // JSON field for additional data
    ProcessedAt     *time.Time       `json:"processed_at,omitempty"`
    ReversedAt      *time.Time       `json:"reversed_at,omitempty"`
    ReversalReason  string           `json:"reversal_reason,omitempty"`
    CreatedAt       time.Time        `json:"created_at"`
    UpdatedAt       time.Time        `json:"updated_at"`
}

// JournalEntry represents a double-entry bookkeeping journal entry
type JournalEntry struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    TransactionID uint      `gorm:"not null" json:"transaction_id"`
    AccountID     uint      `gorm:"not null" json:"account_id"`
    Debit         float64   `gorm:"not null;default:0" json:"debit"`
    Credit        float64   `gorm:"not null;default:0" json:"credit"`
    Description   string    `json:"description,omitempty"`
    CreatedAt     time.Time `json:"created_at"`
}

// Balance represents an account balance at a point in time
type Balance struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    AccountID uint      `gorm:"not null" json:"account_id"`
    Date      time.Time `gorm:"not null" json:"date"`
    Balance   float64   `gorm:"not null" json:"balance"`
    CreatedAt time.Time `json:"created_at"`
}