package card

import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// PushToWalletService handles adding cards to digital wallets (Apple Pay, Google Pay)
type PushToWalletService struct {
    db *gorm.DB
}

// NewPushToWalletService creates a new push to wallet service
func NewPushToWalletService(db *gorm.DB) *PushToWalletService {
    return &PushToWalletService{
        db: db,
    }
}

// DigitalWallet represents a digital wallet (Apple Pay, Google Pay, etc.)
type DigitalWallet struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      uint      `json:"user_id"`
    CardID      string    `json:"card_id"`
    WalletType  string    `json:"wallet_type"` // "apple_pay", "google_pay", etc.
    DeviceID    string    `json:"device_id"`
    Status      string    `json:"status"` // "pending", "active", "inactive"
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// AddToWalletRequest represents a request to add a card to a digital wallet
type AddToWalletRequest struct {
    UserID     uint   `json:"user_id"`
    CardID     string `json:"card_id"`
    WalletType string `json:"wallet_type"`
    DeviceID   string `json:"device_id"`
}

// AddToWalletResponse represents the response from adding a card to a digital wallet
type AddToWalletResponse struct {
    Success     bool   `json:"success"`
    Message     string `json:"message"`
    WalletID    uint   `json:"wallet_id,omitempty"`
    ActivationData string `json:"activation_data,omitempty"`
}

// AddToWallet adds a card to a digital wallet
func (p *PushToWalletService) AddToWallet(ctx context.Context, request AddToWalletRequest) (*AddToWalletResponse, error) {
    // Verify that the card exists and belongs to the user
    var card Card
    err := p.db.Where("id = ? AND cardholder_id IN (SELECT id FROM cardholders WHERE user_id = ?)", request.CardID, request.UserID).First(&card).Error
    if err != nil {
        return &AddToWalletResponse{
            Success: false,
            Message: "Card not found or does not belong to user",
        }, nil
    }
    
    // Check if the card is already added to this wallet type
    var existingWallet DigitalWallet
    err = p.db.Where("card_id = ? AND wallet_type = ? AND device_id = ?", request.CardID, request.WalletType, request.DeviceID).First(&existingWallet).Error
    if err == nil {
        return &AddToWalletResponse{
            Success: false,
            Message: "Card is already added to this wallet",
        }, nil
    }
    
    // Create a digital wallet entry
    digitalWallet := &DigitalWallet{
        UserID:     request.UserID,
        CardID:     request.CardID,
        WalletType: request.WalletType,
        DeviceID:   request.DeviceID,
        Status:     "pending",
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
    }
    
    if err := p.db.Create(digitalWallet).Error; err != nil {
        return &AddToWalletResponse{
            Success: false,
            Message: fmt.Sprintf("Failed to add card to wallet: %v", err),
        }, nil
    }
    
    // Generate activation data (in a real implementation, this would be provided by the card issuer)
    activationData := fmt.Sprintf("activation_data_%d", time.Now().UnixNano())
    
    // Update the digital wallet status to active
    digitalWallet.Status = "active"
    digitalWallet.UpdatedAt = time.Now()
    
    if err := p.db.Save(digitalWallet).Error; err != nil {
        return &AddToWalletResponse{
            Success: false,
            Message: fmt.Sprintf("Failed to activate wallet entry: %v", err),
        }, nil
    }
    
    return &AddToWalletResponse{
        Success:        true,
        Message:        "Card successfully added to wallet",
        WalletID:       digitalWallet.ID,
        ActivationData: activationData,
    }, nil
}

// GetWalletsForUser retrieves all digital wallets for a user
func (p *PushToWalletService) GetWalletsForUser(ctx context.Context, userID uint) ([]DigitalWallet, error) {
    var wallets []DigitalWallet
    err := p.db.Where("user_id = ?", userID).Find(&wallets).Error
    if err != nil {
        return nil, err
    }
    
    return wallets, nil
}

// RemoveFromWallet removes a card from a digital wallet
func (p *PushToWalletService) RemoveFromWallet(ctx context.Context, walletID, userID uint) (*AddToWalletResponse, error) {
    var digitalWallet DigitalWallet
    err := p.db.Where("id = ? AND user_id = ?", walletID, userID).First(&digitalWallet).Error
    if err != nil {
        return &AddToWalletResponse{
            Success: false,
            Message: "Wallet entry not found or does not belong to user",
        }, nil
    }
    
    // Update the digital wallet status to inactive
    digitalWallet.Status = "inactive"
    digitalWallet.UpdatedAt = time.Now()
    
    if err := p.db.Save(&digitalWallet).Error; err != nil {
        return &AddToWalletResponse{
            Success: false,
            Message: fmt.Sprintf("Failed to remove card from wallet: %v", err),
        }, nil
    }
    
    return &AddToWalletResponse{
        Success: true,
        Message: "Card successfully removed from wallet",
    }, nil
}

// GetSupportedWallets returns the list of supported digital wallets
func (p *PushToWalletService) GetSupportedWallets() []string {
    return []string{
        "apple_pay",
        "google_pay",
        "samsung_pay",
    }
}