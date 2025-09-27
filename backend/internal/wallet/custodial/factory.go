package custodial

import "fmt"

// ProviderFactory creates custodial wallet providers
type ProviderFactory struct {
    config map[string]interface{}
}

// NewProviderFactory creates a new provider factory
func NewProviderFactory(config map[string]interface{}) *ProviderFactory {
    return &ProviderFactory{
        config: config,
    }
}

// CreateProvider creates a custodial wallet provider for the specified service
func (f *ProviderFactory) CreateProvider(provider string) (Provider, error) {
    switch provider {
    case "fireblocks":
        apiKey, _ := f.config["fireblocks_api_key"].(string)
        secretKey, _ := f.config["fireblocks_secret_key"].(string)
        baseURL, _ := f.config["fireblocks_base_url"].(string)
        return NewFireblocksProvider(apiKey, secretKey, baseURL), nil
    case "bitgo":
        accessToken, _ := f.config["bitgo_access_token"].(string)
        baseURL, _ := f.config["bitgo_base_url"].(string)
        return NewBitGoProvider(accessToken, baseURL), nil
    case "coinbase":
        apiKey, _ := f.config["coinbase_api_key"].(string)
        secretKey, _ := f.config["coinbase_secret_key"].(string)
        baseURL, _ := f.config["coinbase_base_url"].(string)
        return NewCoinbaseProvider(apiKey, secretKey, baseURL), nil
    default:
        return nil, fmt.Errorf("unsupported custodial wallet provider: %s", provider)
    }
}