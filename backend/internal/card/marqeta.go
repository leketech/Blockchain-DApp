package card

import (
    "context"
    "crypto/rand"
    "fmt"
    "math/big"
    "time"
)

// MarqetaIssuer implements the Issuer interface for Marqeta
type MarqetaIssuer struct {
    apiKey    string
    secretKey string
    baseURL   string
}

// NewMarqetaIssuer creates a new Marqeta issuer
func NewMarqetaIssuer(apiKey, secretKey, baseURL string) *MarqetaIssuer {
    return &MarqetaIssuer{
        apiKey:    apiKey,
        secretKey: secretKey,
        baseURL:   baseURL,
    }
}

// CreateCardholder creates a new cardholder
func (m *MarqetaIssuer) CreateCardholder(ctx context.Context, cardholder *Cardholder) (*Cardholder, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll generate mock values
    
    // Generate a mock cardholder token
    cardholderID := fmt.Sprintf("user-%d", time.Now().UnixNano())
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   cardholder.FirstName,
        LastName:    cardholder.LastName,
        Email:       cardholder.Email,
        Phone:       cardholder.Phone,
        Address:     cardholder.Address,
        DateOfBirth: cardholder.DateOfBirth,
        Status:      "ACTIVE",
        Metadata:    cardholder.Metadata,
    }, nil
}

// GetCardholder retrieves cardholder information
func (m *MarqetaIssuer) GetCardholder(ctx context.Context, cardholderID string) (*Cardholder, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll return a mock cardholder
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   "Jane",
        LastName:    "Smith",
        Email:       "jane.smith@example.com",
        Phone:       "+1987654321",
        Address:     "456 Oak Ave, Town, State 67890",
        DateOfBirth: "1985-05-15",
        Status:      "ACTIVE",
        Metadata: map[string]interface{}{
            "created_at": time.Now().Unix(),
        },
    }, nil
}

// UpdateCardholder updates cardholder information
func (m *MarqetaIssuer) UpdateCardholder(ctx context.Context, cardholderID string, cardholder *Cardholder) (*Cardholder, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll return the updated cardholder
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   cardholder.FirstName,
        LastName:    cardholder.LastName,
        Email:       cardholder.Email,
        Phone:       cardholder.Phone,
        Address:     cardholder.Address,
        DateOfBirth: cardholder.DateOfBirth,
        Status:      "ACTIVE",
        Metadata:    cardholder.Metadata,
    }, nil
}

// CreateCard creates a new card
func (m *MarqetaIssuer) CreateCard(ctx context.Context, cardholderID string, cardType string) (*Card, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll generate mock values
    
    // Generate a mock card token
    cardID := fmt.Sprintf("card-%d", time.Now().UnixNano())
    
    // Generate mock card details
    cardNumber := "4111111111111111"
    if cardType == "virtual" {
        cardNumber = "4000000000000002"
    }
    
    // Generate random expiry date
    expiryMonth := 1 + time.Now().Month()
    expiryYear := time.Now().Year() + 3
    
    // Generate random CVC
    cvc, _ := rand.Int(rand.Reader, big.NewInt(900))
    cvcInt := int(cvc.Int64()) + 100
    
    return &Card{
        ID:              cardID,
        CardholderID:    cardholderID,
        CardNumber:      cardNumber,
        ExpiryMonth:     int(expiryMonth),
        ExpiryYear:      expiryYear,
        CVC:             fmt.Sprintf("%03d", cvcInt),
        Brand:           "visa",
        Type:            cardType,
        Status:          "INACTIVE",
        Currency:        "usd",
        IssuingProvider: "marqeta",
        CreatedAt:       time.Now().Unix(),
    }, nil
}

// GetCard retrieves card information
func (m *MarqetaIssuer) GetCard(ctx context.Context, cardID string) (*Card, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll return a mock card
    
    return &Card{
        ID:              cardID,
        CardholderID:    "user-1234567890",
        CardNumber:      "411111******1111",
        ExpiryMonth:     11,
        ExpiryYear:      2026,
        CVC:             "***",
        Brand:           "visa",
        Type:            "physical",
        Status:          "ACTIVE",
        Currency:        "usd",
        IssuingProvider: "marqeta",
        CreatedAt:       time.Now().Unix() - 172800, // 2 days ago
    }, nil
}

// ActivateCard activates a card
func (m *MarqetaIssuer) ActivateCard(ctx context.Context, cardID string) error {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll just return nil
    return nil
}

// DeactivateCard deactivates a card
func (m *MarqetaIssuer) DeactivateCard(ctx context.Context, cardID string) error {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll just return nil
    return nil
}

// CancelCard cancels a card
func (m *MarqetaIssuer) CancelCard(ctx context.Context, cardID string) error {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll just return nil
    return nil
}

// GetCardTransactions retrieves card transactions
func (m *MarqetaIssuer) GetCardTransactions(ctx context.Context, cardID string, limit, offset int) ([]Transaction, error) {
    // In a real implementation, this would make an API call to Marqeta
    // For demonstration, we'll return mock transactions
    
    transactions := make([]Transaction, 0, limit)
    
    for i := 0; i < limit; i++ {
        transaction := Transaction{
            ID:                fmt.Sprintf("tran-%d-%d", time.Now().UnixNano(), i),
            CardID:            cardID,
            Amount:            float64(i+1) * 45.50,
            Currency:          "usd",
            Status:            "COMPLETION",
            Merchant:          fmt.Sprintf("Store #%d", i+1),
            MerchantCity:      "New York",
            MerchantState:     "NY",
            MerchantCountry:   "US",
            AuthorizationCode: fmt.Sprintf("MC%d", i+1),
            CreatedAt:         time.Now().Unix() - int64(i*7200), // i*2 hours ago
        }
        
        transactions = append(transactions, transaction)
    }
    
    return transactions, nil
}

// GetSupportedCardBrands returns the list of supported card brands
func (m *MarqetaIssuer) GetSupportedCardBrands() []string {
    return []string{
        "visa",
        "mastercard",
    }
}

// GetSupportedCardTypes returns the list of supported card types
func (m *MarqetaIssuer) GetSupportedCardTypes() []string {
    return []string{
        "physical",
        "virtual",
    }
}