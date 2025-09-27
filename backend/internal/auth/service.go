package auth

import (
    "crypto/rand"
    "crypto/subtle"
    "encoding/base64"
    "fmt"
    "time"

    "github.com/blockchain-dapp/backend/internal/pkg/security"

    "golang.org/x/crypto/argon2"
    "gorm.io/gorm"
)

// Service provides authentication operations
type Service struct {
    db *gorm.DB
    jwt *security.JWT
}

// NewService creates a new auth service
func NewService(db *gorm.DB) *Service {
    return &Service{
        db: db,
        jwt: security.NewJWT(security.DefaultJWTConfig()),
    }
}

// PasswordHasher defines the interface for password hashing
type PasswordHasher interface {
    HashPassword(password string) (string, error)
    ComparePassword(hashedPassword, password string) bool
}

// Argon2Hasher implements PasswordHasher using Argon2
type Argon2Hasher struct {
    time    uint32
    memory  uint32
    threads uint8
    keyLen  uint32
}

// NewArgon2Hasher creates a new Argon2 hasher
func NewArgon2Hasher() *Argon2Hasher {
    return &Argon2Hasher{
        time:    1,
        memory:  64 * 1024,
        threads: 4,
        keyLen:  32,
    }
}

// HashPassword hashes a password using Argon2
func (a *Argon2Hasher) HashPassword(password string) (string, error) {
    // Generate a random salt
    salt := make([]byte, 16)
    if _, err := rand.Read(salt); err != nil {
        return "", err
    }

    // Hash the password
    hash := argon2.IDKey([]byte(password), salt, a.time, a.memory, a.threads, a.keyLen)

    // Encode the salt and hash
    b64Salt := base64.StdEncoding.EncodeToString(salt)
    b64Hash := base64.StdEncoding.EncodeToString(hash)

    // Return the formatted hash
    return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.memory, a.time, a.threads, b64Salt, b64Hash), nil
}

// ComparePassword compares a password with a hashed password
func (a *Argon2Hasher) ComparePassword(hashedPassword, password string) bool {
    // Parse the hashed password
    var version int
    var memory uint32
    var time uint32
    var threads uint8
    var salt []byte
    var hash []byte

    if _, err := fmt.Sscanf(hashedPassword, "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", &version, &memory, &time, &threads, &salt, &hash); err != nil {
        return false
    }

    // Decode the salt and hash
    b64Salt, err := base64.StdEncoding.DecodeString(string(salt))
    if err != nil {
        return false
    }

    b64Hash, err := base64.StdEncoding.DecodeString(string(hash))
    if err != nil {
        return false
    }

    // Hash the password with the same parameters
    otherHash := argon2.IDKey([]byte(password), b64Salt, time, memory, threads, uint32(len(b64Hash)))

    // Compare the hashes
    return subtle.ConstantTimeCompare(b64Hash, otherHash) == 1
}

// CreateUser creates a new user
func (s *Service) CreateUser(user *User) error {
    if user.Email == "" || user.Password == "" {
        return fmt.Errorf("email and password are required")
    }

    // Hash the password
    hasher := NewArgon2Hasher()
    hashedPassword, err := hasher.HashPassword(user.Password)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }

    user.Password = hashedPassword

    return s.db.Create(user).Error
}

// GetUserByEmail retrieves a user by email
func (s *Service) GetUserByEmail(email string) (*User, error) {
    var user User
    err := s.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// GetUserByID retrieves a user by ID
func (s *Service) GetUserByID(id uint) (*User, error) {
    var user User
    err := s.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates a user
func (s *Service) UpdateUser(user *User) error {
    return s.db.Save(user).Error
}

// AuthenticateUser authenticates a user by email and password
func (s *Service) AuthenticateUser(email, password string) (*User, error) {
    user, err := s.GetUserByEmail(email)
    if err != nil {
        return nil, fmt.Errorf("invalid email or password")
    }

    hasher := NewArgon2Hasher()
    if !hasher.ComparePassword(user.Password, password) {
        return nil, fmt.Errorf("invalid email or password")
    }

    return user, nil
}

// CreateSession creates a new session for a user
func (s *Service) CreateSession(userID uint, ip, userAgent string) (*Session, error) {
    // Generate a random token
    tokenBytes := make([]byte, 32)
    if _, err := rand.Read(tokenBytes); err != nil {
        return nil, err
    }
    token := base64.URLEncoding.EncodeToString(tokenBytes)

    session := &Session{
        UserID:    userID,
        Token:     token,
        ExpiresAt: time.Now().Add(24 * time.Hour), // 24 hours
        IP:        ip,
        UserAgent: userAgent,
        CreatedAt: time.Now(),
    }

    return session, s.db.Create(session).Error
}

// GetSessionByToken retrieves a session by token
func (s *Service) GetSessionByToken(token string) (*Session, error) {
    var session Session
    err := s.db.Where("token = ?", token).First(&session).Error
    if err != nil {
        return nil, err
    }
    
    // Check if session has expired
    if session.ExpiresAt.Before(time.Now()) {
        return nil, fmt.Errorf("session has expired")
    }
    
    return &session, nil
}

// DeleteSession deletes a session
func (s *Service) DeleteSession(token string) error {
    return s.db.Where("token = ?", token).Delete(&Session{}).Error
}

// CreatePasswordReset creates a new password reset token
func (s *Service) CreatePasswordReset(userID uint) (*PasswordReset, error) {
    // Generate a random token
    tokenBytes := make([]byte, 32)
    if _, err := rand.Read(tokenBytes); err != nil {
        return nil, err
    }
    token := base64.URLEncoding.EncodeToString(tokenBytes)

    reset := &PasswordReset{
        UserID:    userID,
        Token:     token,
        ExpiresAt: time.Now().Add(1 * time.Hour), // 1 hour
        Used:      false,
        CreatedAt: time.Now(),
    }

    return reset, s.db.Create(reset).Error
}

// GetPasswordResetByToken retrieves a password reset by token
func (s *Service) GetPasswordResetByToken(token string) (*PasswordReset, error) {
    var reset PasswordReset
    err := s.db.Where("token = ?", token).First(&reset).Error
    if err != nil {
        return nil, err
    }
    
    // Check if reset has expired
    if reset.ExpiresAt.Before(time.Now()) {
        return nil, fmt.Errorf("password reset token has expired")
    }
    
    // Check if reset has been used
    if reset.Used {
        return nil, fmt.Errorf("password reset token has already been used")
    }
    
    return &reset, nil
}

// UsePasswordReset marks a password reset as used
func (s *Service) UsePasswordReset(token string) error {
    return s.db.Model(&PasswordReset{}).Where("token = ?", token).Update("used", true).Error
}

// UpdateUserPassword updates a user's password
func (s *Service) UpdateUserPassword(userID uint, newPassword string) error {
    // Hash the new password
    hasher := NewArgon2Hasher()
    hashedPassword, err := hasher.HashPassword(newPassword)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }

    // Update the user's password
    return s.db.Model(&User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// GenerateTokens generates JWT tokens for a user
func (s *Service) GenerateTokens(user *User) (string, string, error) {
    // Generate access token
    accessToken, err := s.jwt.GenerateToken(user.ID, user.Email, user.Role, user.FirstName+" "+user.LastName)
    if err != nil {
        return "", "", fmt.Errorf("failed to generate access token: %w", err)
    }

    // Generate refresh token
    refreshToken, err := s.jwt.GenerateRefreshToken(user.ID, user.Email)
    if err != nil {
        return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
    }

    return accessToken, refreshToken, nil
}

// ValidateRefreshToken validates a refresh token and returns a new access token
func (s *Service) ValidateRefreshToken(refreshToken string) (string, error) {
    // Validate the refresh token
    claims, err := s.jwt.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", fmt.Errorf("invalid refresh token: %w", err)
    }

    // Get user information
    user, err := s.GetUserByEmail(claims.Subject)
    if err != nil {
        return "", fmt.Errorf("user not found: %w", err)
    }

    // Generate new access token
    accessToken, err := s.jwt.GenerateToken(user.ID, user.Email, user.Role, user.FirstName+" "+user.LastName)
    if err != nil {
        return "", fmt.Errorf("failed to generate access token: %w", err)
    }

    return accessToken, nil
}