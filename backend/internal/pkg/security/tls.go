package security

import (
    "crypto/tls"
    "fmt"
    "log"
)

// TLSConfig holds TLS configuration
type TLSConfig struct {
    CertFile string
    KeyFile  string
    MinVersion uint16
    MaxVersion uint16
    CipherSuites []uint16
}

// DefaultTLSConfig returns a secure TLS configuration
func DefaultTLSConfig() *TLSConfig {
    return &TLSConfig{
        MinVersion: tls.VersionTLS12,
        MaxVersion: tls.VersionTLS13,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }
}

// LoadTLSConfig loads TLS configuration from files
func (t *TLSConfig) LoadTLSConfig() (*tls.Config, error) {
    if t.CertFile == "" || t.KeyFile == "" {
        return nil, fmt.Errorf("certificate and key files are required")
    }

    cert, err := tls.LoadX509KeyPair(t.CertFile, t.KeyFile)
    if err != nil {
        return nil, fmt.Errorf("failed to load TLS certificate: %w", err)
    }

    config := &tls.Config{
        Certificates: []tls.Certificate{cert},
        MinVersion:   t.MinVersion,
        MaxVersion:   t.MaxVersion,
        CipherSuites: t.CipherSuites,
        PreferServerCipherSuites: true,
    }

    return config, nil
}

// ConfigureTLS configures TLS for the Fiber app
func ConfigureTLS(certFile, keyFile string) *tls.Config {
    config := DefaultTLSConfig()
    config.CertFile = certFile
    config.KeyFile = keyFile

    tlsConfig, err := config.LoadTLSConfig()
    if err != nil {
        log.Fatalf("Failed to configure TLS: %v", err)
    }

    return tlsConfig
}