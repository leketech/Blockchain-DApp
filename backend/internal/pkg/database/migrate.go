package database

import (
    "log"

    "github.com/blockchain-dapp/backend/internal/accounting"
    "github.com/blockchain-dapp/backend/internal/admin"
    "github.com/blockchain-dapp/backend/internal/auth"
    "github.com/blockchain-dapp/backend/internal/card"
    "github.com/blockchain-dapp/backend/internal/kyc"
    "github.com/blockchain-dapp/backend/internal/payments"
    "github.com/blockchain-dapp/backend/internal/wallet"

    "gorm.io/gorm"
)

// Migrate runs the database migrations
func Migrate(db *gorm.DB) {
    // Run migrations for all models
    models := []interface{}{
        // Auth models
        &auth.User{},
        &auth.Session{},
        &auth.PasswordReset{},
        
        // Wallet models
        &wallet.Wallet{},
        &wallet.Transaction{},
        &wallet.CustodialWallet{},
        
        // Payment models
        &payments.PaymentRecord{},
        &payments.CustomerRecord{},
        &payments.PaymentMethodRecord{},
        
        // Card models
        &card.CardRecord{},
        &card.CardholderRecord{},
        &card.TransactionRecord{},
        
        // KYC models
        &kyc.User{},
        &kyc.Document{},
        &kyc.Verification{},
        
        // Accounting models
        &accounting.Account{},
        &accounting.Transaction{},
        &accounting.JournalEntry{},
        &accounting.Balance{},
        
        // Admin models
        &admin.SupportTicket{},
        &admin.AuditLog{},
        &admin.SystemMetric{},
    }

    for _, model := range models {
        if err := db.AutoMigrate(model); err != nil {
            log.Fatalf("Failed to migrate model %T: %v", model, err)
        }
    }

    log.Println("Database migration completed successfully")
}