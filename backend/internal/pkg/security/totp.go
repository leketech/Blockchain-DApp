package security

import (
    "crypto/hmac"
    "crypto/rand"
    "crypto/sha1"
    "encoding/base32"
    "fmt"
    "math"
    "net/url"
    "strconv"
    "strings"
    "time"
)

// TOTPConfig holds TOTP configuration
type TOTPConfig struct {
    Issuer      string
    Algorithm   string
    Digits      int
    Period      int
    SecretSize  int
}

// DefaultTOTPConfig returns a default TOTP configuration
func DefaultTOTPConfig() *TOTPConfig {
    return &TOTPConfig{
        Issuer:     "BlockchainDApp",
        Algorithm:  "SHA1",
        Digits:     6,
        Period:     30,
        SecretSize: 20,
    }
}

// TOTP holds TOTP implementation
type TOTP struct {
    config *TOTPConfig
}

// NewTOTP creates a new TOTP instance
func NewTOTP(config *TOTPConfig) *TOTP {
    if config == nil {
        config = DefaultTOTPConfig()
    }
    return &TOTP{config: config}
}

// GenerateSecret generates a new secret key
func (t *TOTP) GenerateSecret() (string, error) {
    secret := make([]byte, t.config.SecretSize)
    if _, err := rand.Read(secret); err != nil {
        return "", err
    }
    
    return base32.StdEncoding.EncodeToString(secret), nil
}

// GenerateOTP generates a TOTP code
func (t *TOTP) GenerateOTP(secret string, timestamp int64) (string, error) {
    // Decode the secret
    decodedSecret, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
    if err != nil {
        return "", err
    }

    // Calculate the counter value
    counter := timestamp / int64(t.config.Period)

    // Convert counter to bytes
    counterBytes := make([]byte, 8)
    for i := 7; i >= 0; i-- {
        counterBytes[i] = byte(counter & 0xff)
        counter >>= 8
    }

    // Generate HMAC-SHA1
    mac := hmac.New(sha1.New, decodedSecret)
    mac.Write(counterBytes)
    hash := mac.Sum(nil)

    // Dynamic truncation
    offset := hash[len(hash)-1] & 0x0f
    binary := ((uint32(hash[offset]) & 0x7f) << 24) |
        ((uint32(hash[offset+1]) & 0xff) << 16) |
        ((uint32(hash[offset+2]) & 0xff) << 8) |
        (uint32(hash[offset+3]) & 0xff)

    // Generate OTP
    otp := binary % uint32(math.Pow10(t.config.Digits))
    
    // Pad with zeros if necessary
    return fmt.Sprintf("%0*d", t.config.Digits, otp), nil
}

// GetCurrentOTP generates the current TOTP code
func (t *TOTP) GetCurrentOTP(secret string) (string, error) {
    return t.GenerateOTP(secret, time.Now().Unix())
}

// VerifyOTP verifies a TOTP code
func (t *TOTP) VerifyOTP(secret, otp string) bool {
    // Get the current OTP
    currentOTP, err := t.GetCurrentOTP(secret)
    if err != nil {
        return false
    }

    // Check if they match
    return currentOTP == otp
}

// GenerateQRCodeURL generates a QR code URL for TOTP setup
func (t *TOTP) GenerateQRCodeURL(accountName, secret string) string {
    // Create the URL
    u := url.URL{}
    u.Scheme = "otpauth"
    u.Host = "totp"
    u.Path = fmt.Sprintf("%s:%s", t.config.Issuer, accountName)
    
    q := u.Query()
    q.Set("secret", secret)
    q.Set("issuer", t.config.Issuer)
    q.Set("algorithm", t.config.Algorithm)
    q.Set("digits", strconv.Itoa(t.config.Digits))
    q.Set("period", strconv.Itoa(t.config.Period))
    u.RawQuery = q.Encode()
    
    return u.String()
}