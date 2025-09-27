package security

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

// KMSClient represents a KMS client interface
type KMSClient interface {
    Encrypt(plaintext []byte) ([]byte, error)
    Decrypt(ciphertext []byte) ([]byte, error)
    GenerateKey() ([]byte, error)
}

// LocalKMS is a local implementation of KMS for development
type LocalKMS struct {
    key []byte
}

// NewLocalKMS creates a new local KMS instance
func NewLocalKMS(masterKey string) *LocalKMS {
    // In a real implementation, you would use a proper key derivation function
    // For demonstration, we'll just use the master key directly
    key := make([]byte, 32)
    copy(key, masterKey)
    return &LocalKMS{key: key}
}

// Encrypt encrypts data using AES-GCM
func (k *LocalKMS) Encrypt(plaintext []byte) ([]byte, error) {
    // Create a new AES cipher
    block, err := aes.NewCipher(k.key)
    if err != nil {
        return nil, err
    }

    // Create a new GCM cipher
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    // Create a nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    // Encrypt the data
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

    return ciphertext, nil
}

// Decrypt decrypts data using AES-GCM
func (k *LocalKMS) Decrypt(ciphertext []byte) ([]byte, error) {
    // Create a new AES cipher
    block, err := aes.NewCipher(k.key)
    if err != nil {
        return nil, err
    }

    // Create a new GCM cipher
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    // Get the nonce size
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }

    // Extract the nonce and ciphertext
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    // Decrypt the data
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

// GenerateKey generates a new key
func (k *LocalKMS) GenerateKey() ([]byte, error) {
    key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        return nil, err
    }
    return key, nil
}

// EncryptString encrypts a string and returns a base64 encoded string
func (k *LocalKMS) EncryptString(plaintext string) (string, error) {
    ciphertext, err := k.Encrypt([]byte(plaintext))
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptString decrypts a base64 encoded string
func (k *LocalKMS) DecryptString(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    
    plaintext, err := k.Decrypt(data)
    if err != nil {
        return "", err
    }
    
    return string(plaintext), nil
}