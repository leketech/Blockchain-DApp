package card

import (
    "time"

    "gorm.io/gorm"
)

// CardRecord represents a card record in the database
type CardRecord struct {
    ID              uint           `gorm:"primaryKey" json:"id"`
    ExternalID      string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Issuer          string         `gorm:"not null" json:"issuer"` // stripe, marqeta, etc.
    CardholderID    string         `gorm:"not null" json:"cardholder_id"`
    Brand           string         `gorm:"not null" json:"brand"`
    Type            string         `gorm:"not null" json:"type"`
    Status          string         `gorm:"not null" json:"status"`
    Currency        string         `gorm:"not null" json:"currency"`
    Last4           string         `json:"last4"`
    ExpiryMonth     int            `json:"expiry_month"`
    ExpiryYear      int            `json:"expiry_year"`
    Metadata        string         `json:"metadata"` // JSON string
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// CardholderRecord represents a cardholder record in the database
type CardholderRecord struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    ExternalID  string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Issuer      string         `gorm:"not null" json:"issuer"` // stripe, marqeta, etc.
    FirstName   string         `gorm:"not null" json:"first_name"`
    LastName    string         `gorm:"not null" json:"last_name"`
    Email       string         `gorm:"not null" json:"email"`
    Phone       string         `json:"phone"`
    Address     string         `json:"address"`
    DateOfBirth string         `json:"date_of_birth"`
    Status      string         `gorm:"not null" json:"status"`
    Metadata    string         `json:"metadata"` // JSON string
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TransactionRecord represents a card transaction record in the database
type TransactionRecord struct {
    ID                uint           `gorm:"primaryKey" json:"id"`
    ExternalID        string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Issuer            string         `gorm:"not null" json:"issuer"` // stripe, marqeta, etc.
    CardID            string         `gorm:"not null" json:"card_id"`
    Amount            float64        `gorm:"not null" json:"amount"`
    Currency          string         `gorm:"not null" json:"currency"`
    Status            string         `gorm:"not null" json:"status"`
    Merchant          string         `json:"merchant"`
    MerchantCity      string         `json:"merchant_city"`
    MerchantState     string         `json:"merchant_state"`
    MerchantCountry   string         `json:"merchant_country"`
    AuthorizationCode string         `json:"authorization_code"`
    Metadata          string         `json:"metadata"` // JSON string
    CreatedAt         time.Time      `json:"created_at"`
    UpdatedAt         time.Time      `json:"updated_at"`
    DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}