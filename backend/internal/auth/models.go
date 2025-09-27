package auth

import (
    "time"

    "gorm.io/gorm"
)

// User represents a user in the system
type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Email     string         `gorm:"uniqueIndex;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`
    FirstName string         `gorm:"not null" json:"first_name"`
    LastName  string         `gorm:"not null" json:"last_name"`
    Phone     string         `json:"phone"`
    Address   string         `json:"address"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    Role      string         `gorm:"default:'user'" json:"role"`
    MFAEnabled bool          `gorm:"default:false" json:"mfa_enabled"`
    MFASecret string         `gorm:"-" json:"-"` // Not stored in DB for security, but used during setup
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Session represents a user session
type Session struct {
    ID        uint      `gorm:"primaryKey" json:"-"`
    UserID    uint      `gorm:"not null" json:"-"`
    Token     string    `gorm:"uniqueIndex;not null" json:"-"`
    ExpiresAt time.Time `gorm:"not null" json:"-"`
    IP        string    `json:"-"` 
    UserAgent string    `json:"-"`
    CreatedAt time.Time `json:"-"`
}

// PasswordReset represents a password reset token
type PasswordReset struct {
    ID        uint      `gorm:"primaryKey" json:"-"`
    UserID    uint      `gorm:"not null" json:"-"`
    Token     string    `gorm:"uniqueIndex;not null" json:"-"`
    ExpiresAt time.Time `gorm:"not null" json:"-"`
    Used      bool      `gorm:"default:false" json:"-"`
    CreatedAt time.Time `json:"-"`
}

// TableName overrides the table name for User
func (User) TableName() string {
    return "users"
}

// TableName overrides the table name for Session
func (Session) TableName() string {
    return "user_sessions"
}

// TableName overrides the table name for PasswordReset
func (PasswordReset) TableName() string {
    return "password_resets"
}