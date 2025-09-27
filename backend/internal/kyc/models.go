package kyc

import (
    "time"
)

// VerificationStatus represents the status of a KYC verification
type VerificationStatus string

const (
    StatusPending   VerificationStatus = "pending"
    StatusApproved  VerificationStatus = "approved"
    StatusRejected  VerificationStatus = "rejected"
    StatusReview    VerificationStatus = "review"
    StatusCancelled VerificationStatus = "cancelled"
)

// DocumentType represents the type of identification document
type DocumentType string

const (
    DocPassport      DocumentType = "passport"
    DocIDCard        DocumentType = "id_card"
    DocDriverLicense DocumentType = "driver_license"
    DocUtilityBill   DocumentType = "utility_bill"
    DocBankStatement DocumentType = "bank_statement"
)

// User represents a user in the KYC system
type User struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `gorm:"not null" json:"user_id"`
    FirstName   string    `gorm:"not null" json:"first_name"`
    LastName    string    `gorm:"not null" json:"last_name"`
    Email       string    `gorm:"not null" json:"email"`
    Phone       string    `gorm:"not null" json:"phone"`
    DateOfBirth time.Time `gorm:"not null" json:"date_of_birth"`
    Address     string    `gorm:"not null" json:"address"`
    Country     string    `gorm:"not null" json:"country"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Document represents a KYC document
type Document struct {
    ID              uint           `gorm:"primaryKey" json:"id"`
    UserID          uint           `gorm:"not null" json:"user_id"`
    Type            DocumentType   `gorm:"not null" json:"type"`
    FileName        string         `gorm:"not null" json:"file_name"`
    FileURL         string         `gorm:"not null" json:"file_url"`
    Status          string         `gorm:"not null" json:"status"`
    VerifiedBy      *uint          `json:"verified_by,omitempty"`
    VerifiedAt      *time.Time     `json:"verified_at,omitempty"`
    RejectionReason string         `json:"rejection_reason,omitempty"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
}

// Verification represents a KYC verification process
type Verification struct {
    ID              uint              `gorm:"primaryKey" json:"id"`
    UserID          uint              `gorm:"not null" json:"user_id"`
    Status          VerificationStatus `gorm:"not null" json:"status"`
    ProviderID      string            `gorm:"not null" json:"provider_id"`
    ProviderName    string            `gorm:"not null" json:"provider_name"`
    ReferenceID     string            `gorm:"not null" json:"reference_id"`
    RiskScore       float64           `json:"risk_score,omitempty"`
    VerificationURL string            `json:"verification_url,omitempty"`
    CallbackURL     string            `json:"callback_url,omitempty"`
    Metadata        string            `json:"metadata,omitempty"` // JSON field for additional data
    CompletedAt     *time.Time        `json:"completed_at,omitempty"`
    CreatedAt       time.Time         `json:"created_at"`
    UpdatedAt       time.Time         `json:"updated_at"`
}