package card

import "context"

// Card represents a physical or virtual card
type Card struct {
    ID              string
    CardholderID    string
    CardNumber      string
    ExpiryMonth     int
    ExpiryYear      int
    CVC             string
    Brand           string // visa, mastercard, etc.
    Type            string // physical, virtual
    Status          string // active, inactive, cancelled
    Currency        string
    IssuingProvider string // stripe, marqeta, etc.
    CreatedAt       int64
}

// Cardholder represents a cardholder
type Cardholder struct {
    ID          string
    FirstName   string
    LastName    string
    Email       string
    Phone       string
    Address     string
    DateOfBirth string
    Status      string
    Metadata    map[string]interface{}
}

// Transaction represents a card transaction
type Transaction struct {
    ID            string
    CardID        string
    Amount        float64
    Currency      string
    Status        string
    Merchant      string
    MerchantCity  string
    MerchantState string
    MerchantCountry string
    AuthorizationCode string
    CreatedAt     int64
}

// Issuer defines the interface for card issuing providers
type Issuer interface {
    // CreateCardholder creates a new cardholder
    CreateCardholder(ctx context.Context, cardholder *Cardholder) (*Cardholder, error)
    
    // GetCardholder retrieves cardholder information
    GetCardholder(ctx context.Context, cardholderID string) (*Cardholder, error)
    
    // UpdateCardholder updates cardholder information
    UpdateCardholder(ctx context.Context, cardholderID string, cardholder *Cardholder) (*Cardholder, error)
    
    // CreateCard creates a new card
    CreateCard(ctx context.Context, cardholderID string, cardType string) (*Card, error)
    
    // GetCard retrieves card information
    GetCard(ctx context.Context, cardID string) (*Card, error)
    
    // ActivateCard activates a card
    ActivateCard(ctx context.Context, cardID string) error
    
    // DeactivateCard deactivates a card
    DeactivateCard(ctx context.Context, cardID string) error
    
    // CancelCard cancels a card
    CancelCard(ctx context.Context, cardID string) error
    
    // GetCardTransactions retrieves card transactions
    GetCardTransactions(ctx context.Context, cardID string, limit, offset int) ([]Transaction, error)
    
    // GetSupportedCardBrands returns the list of supported card brands
    GetSupportedCardBrands() []string
    
    // GetSupportedCardTypes returns the list of supported card types
    GetSupportedCardTypes() []string
}