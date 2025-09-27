package card

import "fmt"

// IssuerFactory creates card issuing providers
type IssuerFactory struct {
    config map[string]interface{}
}

// NewIssuerFactory creates a new issuer factory
func NewIssuerFactory(config map[string]interface{}) *IssuerFactory {
    return &IssuerFactory{
        config: config,
    }
}

// CreateIssuer creates a card issuing provider for the specified service
func (f *IssuerFactory) CreateIssuer(issuer string) (Issuer, error) {
    switch issuer {
    case "stripe":
        secretKey, _ := f.config["stripe_issuing_secret_key"].(string)
        baseURL, _ := f.config["stripe_issuing_base_url"].(string)
        return NewStripeIssuer(secretKey, baseURL), nil
    case "marqeta":
        apiKey, _ := f.config["marqeta_api_key"].(string)
        secretKey, _ := f.config["marqeta_secret_key"].(string)
        baseURL, _ := f.config["marqeta_base_url"].(string)
        return NewMarqetaIssuer(apiKey, secretKey, baseURL), nil
    default:
        return nil, fmt.Errorf("unsupported card issuing provider: %s", issuer)
    }
}