package payments

import "fmt"

// ProcessorFactory creates payment processors
type ProcessorFactory struct {
    config map[string]interface{}
}

// NewProcessorFactory creates a new processor factory
func NewProcessorFactory(config map[string]interface{}) *ProcessorFactory {
    return &ProcessorFactory{
        config: config,
    }
}

// CreateProcessor creates a payment processor for the specified service
func (f *ProcessorFactory) CreateProcessor(processor string) (Processor, error) {
    switch processor {
    case "stripe":
        secretKey, _ := f.config["stripe_secret_key"].(string)
        baseURL, _ := f.config["stripe_base_url"].(string)
        return NewStripeProcessor(secretKey, baseURL), nil
    case "checkout":
        secretKey, _ := f.config["checkout_secret_key"].(string)
        publicKey, _ := f.config["checkout_public_key"].(string)
        baseURL, _ := f.config["checkout_base_url"].(string)
        return NewCheckoutProcessor(secretKey, publicKey, baseURL), nil
    default:
        return nil, fmt.Errorf("unsupported payment processor: %s", processor)
    }
}