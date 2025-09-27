package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/blockchain-dapp/backend/internal/pkg/config"
    "github.com/blockchain-dapp/backend/internal/pkg/database"
    "github.com/blockchain-dapp/backend/internal/pkg/logger"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
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
    // v1 := app.Group("/api/v1")

    // Wallet routes
    // wallet.SetupRoutes(v1, db)

    // Authentication routes
    // auth.SetupRoutes(v1, db)

    // Payment routes
    // payments.SetupRoutes(v1, db)

    // Card issuing routes
    // card.SetupRoutes(v1, db)

    // KYC/AML routes
    // kyc.SetupRoutes(v1, db)

    // Accounting/Ledger routes
    // accounting.SetupRoutes(v1, db)

    // Admin routes
    // admin.SetupRoutes(v1, db)
}

func main() {
    // Initialize logger
    logger.Init()

    // Load configuration
    cfg := config.Load()

    // Add a small delay to ensure database is ready
    log.Println("Waiting for database to be ready...")
    time.Sleep(5 * time.Second)

    // Connect to database
    db := database.Connect(cfg.DatabaseURL)
    defer database.Close(db)

    // Run database migrations
    database.Migrate(db)

    // Create fiber app
    app := fiber.New(fiber.Config{
        Prefork:       false,
        CaseSensitive: true,
        StrictRouting: true,
        ServerHeader:  "Blockchain-DApp",
        AppName:       "Blockchain DApp Backend v1.0",
    })

    // Middleware
    app.Use(recover.New())
    app.Use(cors.New())
    app.Use(fiberlogger.New())

    // Setup routes
    setupRoutes(app, db)

    // Start server
    log.Printf("🚀 Starting server on port %s", cfg.Port)
    
    // Create channel to listen for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Run server in goroutine
    go func() {
        if err := app.Listen(":" + cfg.Port); err != nil {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()

    // Wait for interrupt signal
    <-quit
    log.Println("Shutting down server...")

    // Graceful shutdown with timeout
    // ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    // defer cancel()
    
    if err := app.Shutdown(); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exiting")
}