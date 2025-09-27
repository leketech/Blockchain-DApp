package kyc

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles KYC HTTP requests
type Handler struct {
    service *Service
}

// NewHandler creates a new KYC handler
func NewHandler(db *gorm.DB) *Handler {
    // TODO: Initialize with a real provider
    return &Handler{
        service: NewService(db, nil),
    }
}

// SetupRoutes sets up the KYC routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    kyc := router.Group("/kyc")
    {
        // Users
        kyc.Post("/users", handler.CreateUser)
        kyc.Get("/users/:id", handler.GetUser)
        
        // Documents
        kyc.Post("/documents", handler.UploadDocument)
        kyc.Get("/users/:userId/documents", handler.GetUserDocuments)
        
        // Verifications
        kyc.Post("/verifications", handler.SubmitVerification)
        kyc.Get("/verifications/:id", handler.GetVerification)
        kyc.Put("/verifications/:id/status", handler.UpdateVerificationStatus)
        kyc.Get("/users/:userId/verifications", handler.GetUserVerifications)
    }
}

// CreateUser creates a new user in the KYC system
func (h *Handler) CreateUser(c *fiber.Ctx) error {
    var user User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdUser, err := h.service.CreateUser(c.Context(), user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create user",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdUser)
}

// GetUser retrieves a user by ID
func (h *Handler) GetUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    user, err := h.service.GetUser(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }

    return c.JSON(user)
}

// UploadDocument uploads a new document for a user
func (h *Handler) UploadDocument(c *fiber.Ctx) error {
    var document Document
    if err := c.BodyParser(&document); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdDocument, err := h.service.UploadDocument(c.Context(), document)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot upload document",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdDocument)
}

// GetUserDocuments retrieves all documents for a user
func (h *Handler) GetUserDocuments(c *fiber.Ctx) error {
    userId, err := strconv.Atoi(c.Params("userId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    documents, err := h.service.GetDocuments(c.Context(), uint(userId))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve documents",
        })
    }

    return c.JSON(documents)
}

// SubmitVerification submits a new verification request
func (h *Handler) SubmitVerification(c *fiber.Ctx) error {
    var request struct {
        UserID uint `json:"user_id"`
    }
    
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    verification, err := h.service.SubmitVerification(c.Context(), request.UserID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot submit verification",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(verification)
}

// GetVerification retrieves a verification by ID
func (h *Handler) GetVerification(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid verification ID",
        })
    }

    verification, err := h.service.GetVerification(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Verification not found",
        })
    }

    return c.JSON(verification)
}

// UpdateVerificationStatus updates a verification's status
func (h *Handler) UpdateVerificationStatus(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid verification ID",
        })
    }

    var request struct {
        Status VerificationStatus `json:"status"`
    }
    
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.UpdateVerificationStatus(c.Context(), uint(id), request.Status); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update verification status",
        })
    }

    return c.JSON(fiber.Map{"message": "Verification status updated successfully"})
}

// GetUserVerifications retrieves verifications for a user
func (h *Handler) GetUserVerifications(c *fiber.Ctx) error {
    userId, err := strconv.Atoi(c.Params("userId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    verifications, err := h.service.GetVerificationsByUser(c.Context(), uint(userId), 10, 0)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve verifications",
        })
    }

    return c.JSON(verifications)
}