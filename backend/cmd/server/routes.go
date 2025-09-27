package main

import (
    "github.com/blockchain-dapp/backend/internal/accounting"
    "github.com/blockchain-dapp/backend/internal/admin"
    "github.com/blockchain-dapp/backend/internal/auth"
    "github.com/blockchain-dapp/backend/internal/card"
    "github.com/blockchain-dapp/backend/internal/kyc"
    "github.com/blockchain-dapp/backend/internal/payments"
    "github.com/blockchain-dapp/backend/internal/wallet"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
    // Health check endpoint
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":  "ok",
            "message": "Blockchain DApp Backend is running!",
        })
    })

    // API v1 routes
    v1 := app.Group("/api/v1")

    // Wallet routes
    wallet.SetupRoutes(v1, db)

    // Authentication routes
    auth.SetupRoutes(v1, db)

    // Payment routes
    payments.SetupRoutes(v1, db)

    // Card issuing routes
    card.SetupRoutes(v1, db)

    // KYC/AML routes
    kyc.SetupRoutes(v1, db)

    // Accounting/Ledger routes
    accounting.SetupRoutes(v1, db)

    // Admin routes
    admin.SetupRoutes(v1, db)
}