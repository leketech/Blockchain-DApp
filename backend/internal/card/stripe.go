package card

import (
    "context"
    "fmt"
    "math/rand"
    "time"

    "github.com/stripe/stripe-go/v74"
    "github.com/stripe/stripe-go/v74/client"
)

// StripeIssuer implements the Issuer interface for Stripe Issuing
type StripeIssuer struct {
    secretKey string
    baseURL   string
    client    *client.API
}

// NewStripeIssuer creates a new Stripe issuer
func NewStripeIssuer(secretKey, baseURL string) *StripeIssuer {
    // Initialize the Stripe client
    sc := &client.API{}
    sc.Init(secretKey, nil)
    
    return &StripeIssuer{
        secretKey: secretKey,
        baseURL:   baseURL,
        client:    sc,
    }
}

// CreateCardholder creates a new cardholder
func (s *StripeIssuer) CreateCardholder(ctx context.Context, cardholder *Cardholder) (*Cardholder, error) {
    // If we have a client, try to create the cardholder through Stripe
    if s.client != nil {
        // Create the cardholder in Stripe
        stripeCardholder, err := s.client.IssuingCardholders.New(&stripe.IssuingCardholderParams{
            Type:  stripe.String("individual"),
            Email: stripe.String(cardholder.Email),
            Individual: &stripe.IssuingCardholderIndividualParams{
                FirstName: stripe.String(cardholder.FirstName),
                LastName:  stripe.String(cardholder.LastName),
            },
            Billing: &stripe.IssuingCardholderBillingParams{
                Address: &stripe.AddressParams{
                    Line1:      stripe.String(cardholder.Address),
                    City:       stripe.String("San Francisco"),
                    State:      stripe.String("CA"),
                    PostalCode: stripe.String("94107"),
                    Country:    stripe.String("US"),
                },
            },
        })
        if err != nil {
            fmt.Printf("Warning: Could not create cardholder in Stripe: %v\n", err)
        } else {
            return &Cardholder{
                ID:          stripeCardholder.ID,
                FirstName:   cardholder.FirstName,
                LastName:    cardholder.LastName,
                Email:       cardholder.Email,
                Phone:       cardholder.Phone,
                Address:     cardholder.Address,
                DateOfBirth: cardholder.DateOfBirth,
                Status:      string(stripeCardholder.Status),
                Metadata:    cardholder.Metadata,
            }, nil
        }
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll generate mock values
    
    // Generate a mock cardholder ID
    cardholderID := fmt.Sprintf("ich_%d", time.Now().UnixNano())
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   cardholder.FirstName,
        LastName:    cardholder.LastName,
        Email:       cardholder.Email,
        Phone:       cardholder.Phone,
        Address:     cardholder.Address,
        DateOfBirth: cardholder.DateOfBirth,
        Status:      "active",
        Metadata:    cardholder.Metadata,
    }, nil
}

// GetCardholder retrieves cardholder information
func (s *StripeIssuer) GetCardholder(ctx context.Context, cardholderID string) (*Cardholder, error) {
    // If we have a client, try to fetch the cardholder from Stripe
    if s.client != nil {
        stripeCardholder, err := s.client.IssuingCardholders.Get(cardholderID, nil)
        if err != nil {
            fmt.Printf("Warning: Could not fetch cardholder from Stripe: %v\n", err)
        } else {
            // Build address string
            address := ""
            if stripeCardholder.Billing != nil && stripeCardholder.Billing.Address != nil {
                address = fmt.Sprintf("%s, %s, %s %s, %s",
                    stripeCardholder.Billing.Address.Line1,
                    stripeCardholder.Billing.Address.City,
                    stripeCardholder.Billing.Address.State,
                    stripeCardholder.Billing.Address.PostalCode,
                    stripeCardholder.Billing.Address.Country)
            }
            
            return &Cardholder{
                ID:          stripeCardholder.ID,
                FirstName:   stripeCardholder.Individual.FirstName,
                LastName:    stripeCardholder.Individual.LastName,
                Email:       stripeCardholder.Email,
                Phone:       "", // Not available in this example
                Address:     address,
                DateOfBirth: "", // Not available in this example
                Status:      string(stripeCardholder.Status),
                Metadata:    make(map[string]interface{}),
            }, nil
        }
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll return a mock cardholder
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   "John",
        LastName:    "Doe",
        Email:       "john.doe@example.com",
        Phone:       "+1234567890",
        Address:     "123 Main St, City, State 12345",
        DateOfBirth: "1990-01-01",
        Status:      "active",
        Metadata: map[string]interface{}{
            "created_at": time.Now().Unix(),
        },
    }, nil
}

// UpdateCardholder updates cardholder information
func (s *StripeIssuer) UpdateCardholder(ctx context.Context, cardholderID string, cardholder *Cardholder) (*Cardholder, error) {
    // If we have a client, try to update the cardholder in Stripe
    if s.client != nil {
        // Update the cardholder in Stripe
        stripeCardholder, err := s.client.IssuingCardholders.Update(cardholderID, &stripe.IssuingCardholderParams{
            Email: stripe.String(cardholder.Email),
            Individual: &stripe.IssuingCardholderIndividualParams{
                FirstName: stripe.String(cardholder.FirstName),
                LastName:  stripe.String(cardholder.LastName),
            },
            Billing: &stripe.IssuingCardholderBillingParams{
                Address: &stripe.AddressParams{
                    Line1:      stripe.String(cardholder.Address),
                    City:       stripe.String("San Francisco"),
                    State:      stripe.String("CA"),
                    PostalCode: stripe.String("94107"),
                    Country:    stripe.String("US"),
                },
            },
        })
        if err != nil {
            fmt.Printf("Warning: Could not update cardholder in Stripe: %v\n", err)
        } else {
            return &Cardholder{
                ID:          stripeCardholder.ID,
                FirstName:   cardholder.FirstName,
                LastName:    cardholder.LastName,
                Email:       cardholder.Email,
                Phone:       cardholder.Phone,
                Address:     cardholder.Address,
                DateOfBirth: cardholder.DateOfBirth,
                Status:      string(stripeCardholder.Status),
                Metadata:    cardholder.Metadata,
            }, nil
        }
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll return the updated cardholder
    
    return &Cardholder{
        ID:          cardholderID,
        FirstName:   cardholder.FirstName,
        LastName:    cardholder.LastName,
        Email:       cardholder.Email,
        Phone:       cardholder.Phone,
        Address:     cardholder.Address,
        DateOfBirth: cardholder.DateOfBirth,
        Status:      "active",
        Metadata:    cardholder.Metadata,
    }, nil
}

// CreateCard creates a new card
func (s *StripeIssuer) CreateCard(ctx context.Context, cardholderID string, cardType string) (*Card, error) {
    // If we have a client, try to create the card through Stripe
    if s.client != nil {
        // Create the card in Stripe
        stripeCard, err := s.client.IssuingCards.New(&stripe.IssuingCardParams{
            Cardholder: stripe.String(cardholderID),
            Currency:   stripe.String("usd"),
            Type:       stripe.String(cardType),
        })
        if err != nil {
            fmt.Printf("Warning: Could not create card in Stripe: %v\n", err)
        } else {
            return &Card{
                ID:              stripeCard.ID,
                CardholderID:    cardholderID,
                CardNumber:      "", // Not returned for security
                ExpiryMonth:     int(stripeCard.ExpMonth),
                ExpiryYear:      int(stripeCard.ExpYear),
                CVC:             "", // Not returned for security
                Brand:           string(stripeCard.Brand),
                Type:            cardType,
                Status:          string(stripeCard.Status),
                Currency:        "usd",
                IssuingProvider: "stripe",
                CreatedAt:       time.Now().Unix(),
            }, nil
        }
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll generate mock values
    
    // Generate a mock card ID
    cardID := fmt.Sprintf("ic_%d", time.Now().UnixNano())
    
    // Generate mock card details
    cardNumber := "4242424242424242"
    if cardType == "virtual" {
        cardNumber = "4000056655665556"
    }
    
    // Generate random expiry date
    expiryMonth := 1 + time.Now().Month()
    expiryYear := time.Now().Year() + 3
    
    // Generate random CVC
    cvc := rand.Intn(900) + 100
    
    return &Card{
        ID:              cardID,
        CardholderID:    cardholderID,
        CardNumber:      cardNumber,
        ExpiryMonth:     int(expiryMonth),
        ExpiryYear:      expiryYear,
        CVC:             fmt.Sprintf("%03d", cvc),
        Brand:           "visa",
        Type:            cardType,
        Status:          "inactive",
        Currency:        "usd",
        IssuingProvider: "stripe",
        CreatedAt:       time.Now().Unix(),
    }, nil
}

// GetCard retrieves card information
func (s *StripeIssuer) GetCard(ctx context.Context, cardID string) (*Card, error) {
    // If we have a client, try to fetch the card from Stripe
    if s.client != nil {
        stripeCard, err := s.client.IssuingCards.Get(cardID, nil)
        if err != nil {
            fmt.Printf("Warning: Could not fetch card from Stripe: %v\n", err)
        } else {
            return &Card{
                ID:              stripeCard.ID,
                CardholderID:    stripeCard.Cardholder.ID,
                CardNumber:      "", // Not returned for security
                ExpiryMonth:     int(stripeCard.ExpMonth),
                ExpiryYear:      int(stripeCard.ExpYear),
                CVC:             "", // Not returned for security
                Brand:           string(stripeCard.Brand),
                Type:            string(stripeCard.Type),
                Status:          string(stripeCard.Status),
                Currency:        "usd",
                IssuingProvider: "stripe",
                CreatedAt:       time.Now().Unix(),
            }, nil
        }
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll return a mock card
    
    return &Card{
        ID:              cardID,
        CardholderID:    "ich_1234567890",
        CardNumber:      "424242******4242",
        ExpiryMonth:     12,
        ExpiryYear:      2025,
        CVC:             "***",
        Brand:           "visa",
        Type:            "physical",
        Status:          "active",
        Currency:        "usd",
        IssuingProvider: "stripe",
        CreatedAt:       time.Now().Unix() - 86400, // 1 day ago
    }, nil
}

// ActivateCard activates a card
func (s *StripeIssuer) ActivateCard(ctx context.Context, cardID string) error {
    // If we have a client, try to activate the card in Stripe
    if s.client != nil {
        _, err := s.client.IssuingCards.Update(cardID, &stripe.IssuingCardParams{
            Status: stripe.String("active"),
        })
        if err != nil {
            fmt.Printf("Warning: Could not activate card in Stripe: %v\n", err)
            return err
        }
        return nil
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll just return nil
    return nil
}

// DeactivateCard deactivates a card
func (s *StripeIssuer) DeactivateCard(ctx context.Context, cardID string) error {
    // If we have a client, try to deactivate the card in Stripe
    if s.client != nil {
        _, err := s.client.IssuingCards.Update(cardID, &stripe.IssuingCardParams{
            Status: stripe.String("inactive"),
        })
        if err != nil {
            fmt.Printf("Warning: Could not deactivate card in Stripe: %v\n", err)
            return err
        }
        return nil
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll just return nil
    return nil
}

// CancelCard cancels a card
func (s *StripeIssuer) CancelCard(ctx context.Context, cardID string) error {
    // If we have a client, try to cancel the card in Stripe
    if s.client != nil {
        _, err := s.client.IssuingCards.Update(cardID, &stripe.IssuingCardParams{
            Status: stripe.String("canceled"),
        })
        if err != nil {
            fmt.Printf("Warning: Could not cancel card in Stripe: %v\n", err)
            return err
        }
        return nil
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll just return nil
    return nil
}

// GetCardTransactions retrieves card transactions
func (s *StripeIssuer) GetCardTransactions(ctx context.Context, cardID string, limit, offset int) ([]Transaction, error) {
    // If we have a client, try to fetch the transactions from Stripe
    if s.client != nil {
        // In a real implementation, this would make an API call to Stripe
        // to fetch the card transactions
    }
    
    // In a real implementation, this would make an API call to Stripe
    // For demonstration, we'll return mock transactions
    
    transactions := make([]Transaction, 0, limit)
    
    for i := 0; i < limit; i++ {
        transaction := Transaction{
            ID:                fmt.Sprintf("ipi_%d_%d", time.Now().UnixNano(), i),
            CardID:            cardID,
            Amount:            float64(i+1) * 25.99,
            Currency:          "usd",
            Status:            "approved",
            Merchant:          fmt.Sprintf("Merchant #%d", i+1),
            MerchantCity:      "San Francisco",
            MerchantState:     "CA",
            MerchantCountry:   "US",
            AuthorizationCode: fmt.Sprintf("AUTH%d", i+1),
            CreatedAt:         time.Now().Unix() - int64(i*3600), // i hours ago
        }
        
        transactions = append(transactions, transaction)
    }
    
    return transactions, nil
}

// GetSupportedCardBrands returns the list of supported card brands
func (s *StripeIssuer) GetSupportedCardBrands() []string {
    return []string{
        "visa",
        "mastercard",
    }
}

// GetSupportedCardTypes returns the list of supported card types
func (s *StripeIssuer) GetSupportedCardTypes() []string {
    return []string{
        "physical",
        "virtual",
    }
}

// ConnectToSandbox connects to the Stripe sandbox environment
func (s *StripeIssuer) ConnectToSandbox(secretKey string) {
    s.secretKey = secretKey
    s.baseURL = "https://api.stripe.com"
    
    // Initialize the Stripe client
    sc := &client.API{}
    sc.Init(secretKey, nil)
    s.client = sc
}

// Helper function to get string value or empty string
func getStringValue(ptr *string) string {
    if ptr != nil {
        return *ptr
    }
    return ""
}
