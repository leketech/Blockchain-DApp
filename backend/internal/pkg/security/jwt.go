package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey     string
	RefreshSecret string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
}

// DefaultJWTConfig returns a default JWT configuration
func DefaultJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:     "your-secret-key-change-in-production",
		RefreshSecret: "your-refresh-secret-change-in-production",
		TokenExpiry:   15 * time.Minute,  // 15 minutes for access token
		RefreshExpiry: 7 * 24 * time.Hour, // 7 days for refresh token
	}
}

// JWT holds JWT implementation
type JWT struct {
	config *JWTConfig
}

// Claims represents JWT claims
type Claims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

// NewJWT creates a new JWT instance
func NewJWT(config *JWTConfig) *JWT {
	if config == nil {
		config = DefaultJWTConfig()
	}
	return &JWT{config: config}
}

// GenerateToken generates a new JWT token
func (j *JWT) GenerateToken(userID uint, email, role, fullName string) (string, error) {
	// Create the claims
	claims := &Claims{
		UserID:   userID,
		Email:    email,
		Role:     role,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   email,
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString([]byte(j.config.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// GenerateRefreshToken generates a new refresh token
func (j *JWT) GenerateRefreshToken(userID uint, email string) (string, error) {
	// Create the claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString([]byte(j.config.RefreshSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates a JWT token
func (j *JWT) ValidateToken(tokenString string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token
func (j *JWT) ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.config.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}