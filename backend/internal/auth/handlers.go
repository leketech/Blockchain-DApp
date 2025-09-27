package auth

import (
    "fmt"

    "github.com/blockchain-dapp/backend/internal/pkg/security"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles authentication HTTP requests
type Handler struct {
    service *Service
}

// NewHandler creates a new auth handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewService(db),
    }
}

// SetupRoutes sets up the authentication routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    auth := router.Group("/auth")
    {
        auth.Post("/register", handler.Register)
        auth.Post("/login", handler.Login)
        auth.Post("/logout", handler.Logout)
        auth.Post("/password/reset", handler.RequestPasswordReset)
        auth.Post("/password/reset/confirm", handler.ConfirmPasswordReset)
        auth.Post("/refresh", handler.RefreshToken)
        auth.Post("/mfa/setup", handler.SetupMFA) // New MFA setup endpoint
        auth.Post("/mfa/verify", handler.VerifyMFA) // New MFA verification endpoint
    }
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=8"`
    FirstName string `json:"first_name" validate:"required"`
    LastName  string `json:"last_name" validate:"required"`
    Phone     string `json:"phone"`
    Address   string `json:"address"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

// RefreshTokenRequest represents the request body for token refresh
type RefreshTokenRequest struct {
    RefreshToken string `json:"refresh_token" validate:"required"`
}

// PasswordResetRequest represents the request body for password reset
type PasswordResetRequest struct {
    Email string `json:"email" validate:"required,email"`
}

// PasswordResetConfirmRequest represents the request body for password reset confirmation
type PasswordResetConfirmRequest struct {
    Token    string `json:"token" validate:"required"`
    Password string `json:"password" validate:"required,min=8"`
}

// MFASetupRequest represents the request body for MFA setup
type MFASetupRequest struct {
    UserID uint `json:"user_id" validate:"required"`
}

// MFAVerifyRequest represents the request body for MFA verification
type MFAVerifyRequest struct {
    UserID   uint   `json:"user_id" validate:"required"`
    OTP      string `json:"otp" validate:"required,len=6"`
    Secret   string `json:"secret" validate:"required"`
}

// Register creates a new user
func (h *Handler) Register(c *fiber.Ctx) error {
    var req RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Create the user
    user := &User{
        Email:     req.Email,
        Password:  req.Password,
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Phone:     req.Phone,
        Address:   req.Address,
        IsActive:  true,
        Role:      "user",
        // MFA is not enabled by default, user needs to set it up separately
        MFAEnabled: false,
    }

    if err := h.service.CreateUser(user); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": fmt.Sprintf("Cannot create user: %v", err),
        })
    }

    // Remove password from response
    user.Password = ""

    return c.Status(fiber.StatusCreated).JSON(user)
}

// Login authenticates a user
func (h *Handler) Login(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Authenticate the user
    user, err := h.service.AuthenticateUser(req.Email, req.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    // Check if MFA is enabled for this user
    if user.MFAEnabled {
        // Return a challenge to the client to request MFA code
        return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
            "message": "MFA required",
            "user_id": user.ID,
            "mfa_required": true,
        })
    }

    // Get IP and user agent
    ip := c.IP()
    userAgent := string(c.Request().Header.UserAgent())

    // Create a session
    session, err := h.service.CreateSession(user.ID, ip, userAgent)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create session",
        })
    }

    // Generate JWT tokens
    accessToken, refreshToken, err := h.service.GenerateTokens(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot generate tokens",
        })
    }

    // Remove password from response
    user.Password = ""

    return c.JSON(fiber.Map{
        "user":         user,
        "access_token":  accessToken,
        "refresh_token": refreshToken,
        "session_token": session.Token,
        "expires":       session.ExpiresAt,
    })
}

// RefreshToken generates a new access token using a refresh token
func (h *Handler) RefreshToken(c *fiber.Ctx) error {
    var req RefreshTokenRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Validate refresh token and generate new access token
    accessToken, err := h.service.ValidateRefreshToken(req.RefreshToken)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid refresh token",
        })
    }

    return c.JSON(fiber.Map{
        "access_token": accessToken,
    })
}

// Logout logs out a user
func (h *Handler) Logout(c *fiber.Ctx) error {
    // Get the token from the Authorization header
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Missing authorization token",
        })
    }

    // Remove "Bearer " prefix if present
    if len(token) > 7 && token[:7] == "Bearer " {
        token = token[7:]
    }

    // Delete the session
    if err := h.service.DeleteSession(token); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot delete session",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Logged out successfully",
    })
}

// RequestPasswordReset requests a password reset
func (h *Handler) RequestPasswordReset(c *fiber.Ctx) error {
    var req PasswordResetRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Get the user by email
    user, err := h.service.GetUserByEmail(req.Email)
    if err != nil {
        // Don't reveal if the email exists or not
        return c.JSON(fiber.Map{
            "message": "If the email exists, a password reset link has been sent",
        })
    }

    // Create a password reset token
    reset, err := h.service.CreatePasswordReset(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create password reset token",
        })
    }

    // In a real implementation, you would send an email with the reset token
    // For now, we'll just return the token in the response
    return c.JSON(fiber.Map{
        "message": "Password reset token created",
        "token":   reset.Token,
    })
}

// ConfirmPasswordReset confirms a password reset
func (h *Handler) ConfirmPasswordReset(c *fiber.Ctx) error {
    var req PasswordResetConfirmRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Get the password reset by token
    reset, err := h.service.GetPasswordResetByToken(req.Token)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid or expired password reset token",
        })
    }

    // Update the user's password
    if err := h.service.UpdateUserPassword(reset.UserID, req.Password); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update password",
        })
    }

    // Mark the reset as used
    if err := h.service.UsePasswordReset(req.Token); err != nil {
        // Log the error but don't return it to the user
        fmt.Printf("Error marking password reset as used: %v\n", err)
    }

    return c.JSON(fiber.Map{
        "message": "Password updated successfully",
    })
}

// SetupMFA sets up MFA for a user
func (h *Handler) SetupMFA(c *fiber.Ctx) error {
    var req MFASetupRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Initialize TOTP service
    totp := security.NewTOTP(security.DefaultTOTPConfig())

    // Generate a new secret for the user
    secret, err := totp.GenerateSecret()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot generate MFA secret",
        })
    }

    // Generate QR code URL for the user to scan
    user, err := h.service.GetUserByID(req.UserID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    qrCodeURL := totp.GenerateQRCodeURL(user.Email, secret)

    return c.JSON(fiber.Map{
        "secret":     secret,
        "qr_code_url": qrCodeURL,
        "message":    "Scan the QR code with your authenticator app",
    })
}

// VerifyMFA verifies an MFA code
func (h *Handler) VerifyMFA(c *fiber.Ctx) error {
    var req MFAVerifyRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Initialize TOTP service
    totp := security.NewTOTP(security.DefaultTOTPConfig())

    // Verify the OTP
    if !totp.VerifyOTP(req.Secret, req.OTP) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid MFA code",
        })
    }

    // Enable MFA for the user
    user, err := h.service.GetUserByID(req.UserID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    user.MFAEnabled = true
    user.MFASecret = req.Secret

    if err := h.service.UpdateUser(user); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot enable MFA for user",
        })
    }

    // Now proceed with login
    // Get IP and user agent
    ip := c.IP()
    userAgent := string(c.Request().Header.UserAgent())

    // Create a session
    session, err := h.service.CreateSession(user.ID, ip, userAgent)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create session",
        })
    }

    // Generate JWT tokens
    accessToken, refreshToken, err := h.service.GenerateTokens(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot generate tokens",
        })
    }

    // Remove password from response
    user.Password = ""

    return c.JSON(fiber.Map{
        "user":         user,
        "access_token":  accessToken,
        "refresh_token": refreshToken,
        "session_token": session.Token,
        "expires":       session.ExpiresAt,
        "message":       "MFA verified and login successful",
    })
}