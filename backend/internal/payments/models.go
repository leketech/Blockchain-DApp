package payments

import (
    "time"

    "gorm.io/gorm"
)

// PaymentRecord represents a payment record in the database
type PaymentRecord struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    ExternalID    string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Processor     string         `gorm:"not null" json:"processor"` // stripe, checkout, etc.
    Amount        float64        `gorm:"not null" json:"amount"`
    Currency      string         `gorm:"not null" json:"currency"`
    Status        string         `gorm:"not null" json:"status"`
    CustomerID    string         `gorm:"not null" json:"customer_id"`
    PaymentMethod string         `gorm:"not null" json:"payment_method"`
    Description   string         `json:"description"`
    Metadata      string         `json:"metadata"` // JSON string
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// CustomerRecord represents a customer record in the database
type CustomerRecord struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    ExternalID string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Processor string         `gorm:"not null" json:"processor"` // stripe, checkout, etc.
    Email     string         `gorm:"not null" json:"email"`
    Name      string         `json:"name"`
    Phone     string         `json:"phone"`
    Address   string         `json:"address"`
    Metadata  string         `json:"metadata"` // JSON string
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// PaymentMethodRecord represents a payment method record in the database
type PaymentMethodRecord struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    ExternalID string         `gorm:"not null;uniqueIndex" json:"external_id"`
    Processor string         `gorm:"not null" json:"processor"` // stripe, checkout, etc.
    Type      string         `gorm:"not null" json:"type"`
    Customer  string         `gorm:"not null" json:"customer"`
    Metadata  string         `json:"metadata"` // JSON string
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}