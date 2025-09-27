package wallet

import (
    "time"

    "gorm.io/gorm"
)

// Wallet represents a user's wallet
type Wallet struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    UserID    uint           `gorm:"not null" json:"user_id"`
    Address   string         `gorm:"not null;uniqueIndex" json:"address"`
    Chain     string         `gorm:"not null" json:"chain"` // bitcoin, ethereum, solana, tron, bnb
    PublicKey string         `gorm:"not null" json:"public_key"`
    Balance   float64        `gorm:"default:0" json:"balance"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Transaction represents a blockchain transaction
type Transaction struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    WalletID      uint           `gorm:"not null" json:"wallet_id"`
    TxHash        string         `gorm:"not null;uniqueIndex" json:"tx_hash"`
    FromAddress   string         `gorm:"not null" json:"from_address"`
    ToAddress     string         `gorm:"not null" json:"to_address"`
    Amount        float64        `gorm:"not null" json:"amount"`
    Chain         string         `gorm:"not null" json:"chain"`
    Status        string         `gorm:"not null" json:"status"` // pending, confirmed, failed
    Confirmations int            `gorm:"default:0" json:"confirmations"`
    GasPrice      float64        `json:"gas_price,omitempty"`
    GasLimit      float64        `json:"gas_limit,omitempty"`
    GasUsed       float64        `json:"gas_used,omitempty"`
    Fee           float64        `json:"fee,omitempty"`
    Memo          string         `json:"memo,omitempty"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// CustodialWallet represents a custodial wallet managed by third-party services
type CustodialWallet struct {
    ID              uint           `gorm:"primaryKey" json:"id"`
    UserID          uint           `gorm:"not null" json:"user_id"`
    ExternalID      string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Provider        string         `gorm:"not null" json:"provider"` // fireblocks, bitgo, coinbase
    Chain           string         `gorm:"not null" json:"chain"`
    Address         string         `gorm:"not null" json:"address"`
    Balance         float64        `gorm:"default:0" json:"balance"`
    IsActive        bool           `gorm:"default:true" json:"is_active"`
    LastSyncedAt    time.Time      `json:"last_synced_at"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}