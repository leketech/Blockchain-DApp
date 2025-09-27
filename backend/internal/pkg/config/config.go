package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
    Port        string
    DatabaseURL string
    JWTSecret   string
    StripeKey   string
    Environment string
}

// Load reads configuration from environment variables
func Load() *Config {
    // Load .env file if it exists
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    return &Config{
        Port:        getEnv("PORT", "8080"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        JWTSecret:   getEnv("JWT_SECRET", ""),
        StripeKey:   getEnv("STRIPE_KEY", ""),
        Environment: getEnv("ENVIRONMENT", "development"),
    }
}

// getEnv returns the value of the environment variable or a default value
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}