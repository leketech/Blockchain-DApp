package database

import (
    "log"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Connect establishes a connection to the PostgreSQL database with retry mechanism
func Connect(dsn string) *gorm.DB {
    var db *gorm.DB
    var err error
    
    // Retry connection for up to 30 seconds
    for i := 0; i < 30; i++ {
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        
        log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
        log.Println("Retrying in 1 second...")
        time.Sleep(1 * time.Second)
    }
    
    if err != nil {
        log.Fatal("Failed to connect to database after 30 attempts:", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Failed to get database instance:", err)
    }

    // Configure connection pool
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)

    log.Println("Connected to database successfully")
    return db
}

// Close closes the database connection
func Close(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        log.Println("Failed to get database instance:", err)
        return
    }
    
    err = sqlDB.Close()
    if err != nil {
        log.Println("Failed to close database connection:", err)
        return
    }
    
    log.Println("Database connection closed")
}