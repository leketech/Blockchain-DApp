package payments

import "context"

// Payment represents a payment transaction
type Payment struct {
    ID            string
    Amount        float64
    Currency      string
    Status        string
    PaymentMethod string
    CustomerID    string
    Description   string
    CreatedAt     int64
}

// Customer represents a customer
type Customer struct {
    ID       string
    Email    string
    Name     string
    Phone    string
    Address  string
    Metadata map[string]interface{}
}

// PaymentMethod represents a payment method
type PaymentMethod struct {
    ID       string
    Type     string // card, bank_account, etc.
    Customer string
    Metadata map[string]interface{}
}

// Processor defines the interface for payment processors
type Processor interface {
    // CreateCustomer creates a new customer
    CreateCustomer(ctx context.Context, customer *Customer) (*Customer, error)
    
    // GetCustomer retrieves customer information
    GetCustomer(ctx context.Context, customerID string) (*Customer, error)
    
    // UpdateCustomer updates customer information
    UpdateCustomer(ctx context.Context, customerID string, customer *Customer) (*Customer, error)
    
    // CreatePaymentMethod creates a new payment method
    CreatePaymentMethod(ctx context.Context, method *PaymentMethod) (*PaymentMethod, error)
    
    // GetPaymentMethod retrieves payment method information
    GetPaymentMethod(ctx context.Context, methodID string) (*PaymentMethod, error)
    
    // Charge creates a new payment charge
    Charge(ctx context.Context, customerID, paymentMethodID string, amount float64, currency, description string) (*Payment, error)
    
    // Refund refunds a payment
    Refund(ctx context.Context, paymentID string, amount float64) (*Payment, error)
    
    // GetPayment retrieves payment information
    GetPayment(ctx context.Context, paymentID string) (*Payment, error)
    
    // ListPayments lists payments for a customer
    ListPayments(ctx context.Context, customerID string, limit, offset int) ([]Payment, error)
    
    // GetSupportedCurrencies returns the list of supported currencies
    GetSupportedCurrencies() []string
}