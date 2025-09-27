package kyc

import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// Provider defines the interface for KYC/AML providers
type Provider interface {
    // SubmitVerification submits a new verification request
    SubmitVerification(ctx context.Context, user User, documents []Document) (*Verification, error)
    
    // GetVerificationStatus retrieves the status of a verification
    GetVerificationStatus(ctx context.Context, verificationID string) (*Verification, error)
    
    // GetSupportedCountries returns the list of supported countries
    GetSupportedCountries() []string
    
    // GetDocumentRequirements returns document requirements for a country
    GetDocumentRequirements(country string) []DocumentType
}

// Service handles KYC/AML operations
type Service struct {
    db       *gorm.DB
    provider Provider
}

// NewService creates a new KYC service
func NewService(db *gorm.DB, provider Provider) *Service {
    return &Service{
        db:       db,
        provider: provider,
    }
}

// InitializeDB initializes the database tables for KYC service
func (s *Service) InitializeDB() error {
    return s.db.AutoMigrate(&User{}, &Document{}, &Verification{})
}

// CreateUser creates a new user in the KYC system
func (s *Service) CreateUser(ctx context.Context, user User) (*User, error) {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return &user, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(ctx context.Context, userID uint) (*User, error) {
    var user User
    if err := s.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return &user, nil
}

// UploadDocument uploads a new document for a user
func (s *Service) UploadDocument(ctx context.Context, document Document) (*Document, error) {
    document.CreatedAt = time.Now()
    document.UpdatedAt = time.Now()
    document.Status = "uploaded"
    
    if err := s.db.WithContext(ctx).Create(&document).Error; err != nil {
        return nil, fmt.Errorf("failed to upload document: %w", err)
    }
    
    return &document, nil
}

// GetDocuments retrieves all documents for a user
func (s *Service) GetDocuments(ctx context.Context, userID uint) ([]Document, error) {
    var documents []Document
    if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Find(&documents).Error; err != nil {
        return nil, fmt.Errorf("failed to get documents: %w", err)
    }
    
    return documents, nil
}

// SubmitVerification submits a new verification request to the provider
func (s *Service) SubmitVerification(ctx context.Context, userID uint) (*Verification, error) {
    // Get user details
    user, err := s.GetUser(ctx, userID)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    // Get user documents
    documents, err := s.GetDocuments(ctx, userID)
    if err != nil {
        return nil, fmt.Errorf("failed to get documents: %w", err)
    }
    
    // Submit to provider
    verification, err := s.provider.SubmitVerification(ctx, *user, documents)
    if err != nil {
        return nil, fmt.Errorf("failed to submit verification: %w", err)
    }
    
    // Save to database
    verification.CreatedAt = time.Now()
    verification.UpdatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(verification).Error; err != nil {
        return nil, fmt.Errorf("failed to save verification: %w", err)
    }
    
    return verification, nil
}

// GetVerification retrieves a verification by ID
func (s *Service) GetVerification(ctx context.Context, verificationID uint) (*Verification, error) {
    var verification Verification
    if err := s.db.WithContext(ctx).First(&verification, "id = ?", verificationID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("verification not found")
        }
        return nil, fmt.Errorf("failed to get verification: %w", err)
    }
    
    return &verification, nil
}

// UpdateVerificationStatus updates a verification's status
func (s *Service) UpdateVerificationStatus(ctx context.Context, verificationID uint, status VerificationStatus) error {
    return s.db.WithContext(ctx).Model(&Verification{}).Where("id = ?", verificationID).Update("status", status).Error
}

// GetVerificationsByUser retrieves verifications for a user
func (s *Service) GetVerificationsByUser(ctx context.Context, userID uint, limit, offset int) ([]Verification, error) {
    var verifications []Verification
    if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&verifications).Error; err != nil {
        return nil, fmt.Errorf("failed to get verifications: %w", err)
    }
    
    return verifications, nil
}