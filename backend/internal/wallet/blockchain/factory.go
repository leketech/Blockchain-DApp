package blockchain

import "fmt"

// AdapterFactory creates blockchain adapters
type AdapterFactory struct {
    config map[string]interface{}
}

// NewAdapterFactory creates a new adapter factory
func NewAdapterFactory(config map[string]interface{}) *AdapterFactory {
    return &AdapterFactory{
        config: config,
    }
}

// CreateAdapter creates a blockchain adapter for the specified chain
func (f *AdapterFactory) CreateAdapter(chain string) (Adapter, error) {
    switch chain {
    case "bitcoin":
        isTestnet, _ := f.config["bitcoin_testnet"].(bool)
        return NewBitcoinAdapter(isTestnet), nil
    case "ethereum":
        rpcURL, _ := f.config["ethereum_rpc_url"].(string)
        return NewEthereumAdapter(rpcURL), nil
    case "solana":
        rpcURL, _ := f.config["solana_rpc_url"].(string)
        return NewSolanaAdapter(rpcURL), nil
    case "tron":
        grpcURL, _ := f.config["tron_grpc_url"].(string)
        return NewTronAdapter(grpcURL), nil
    case "bnb":
        rpcURL, _ := f.config["bnb_rpc_url"].(string)
        return NewBNBadapter(rpcURL), nil
    default:
        return nil, fmt.Errorf("unsupported blockchain: %s", chain)
    }
}