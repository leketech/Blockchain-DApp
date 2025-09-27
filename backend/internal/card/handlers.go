package card

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles card HTTP requests
type Handler struct {
    service *Service
}

// NewHandler creates a new card handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewService(db),
    }
}

// SetupRoutes sets up the card routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    cards := router.Group("/cards")
    {
        cards.Post("/", handler.CreateCard)
        cards.Get("/", handler.GetCards)
        cards.Get("/:id", handler.GetCard)
        cards.Put("/:id", handler.UpdateCard)
        cards.Post("/:id/activate", handler.ActivateCard)
        cards.Post("/:id/deactivate", handler.DeactivateCard)
        cards.Post("/:id/cancel", handler.CancelCard)
        cards.Get("/:id/transactions", handler.GetCardTransactions)
        
        // Cardholders
        cards.Post("/cardholders", handler.CreateCardholder)
        cards.Get("/cardholders", handler.GetCardholders)
        cards.Get("/cardholders/:id", handler.GetCardholder)
        cards.Put("/cardholders/:id", handler.UpdateCardholder)
    }
}

// CreateCard creates a new card
func (h *Handler) CreateCard(c *fiber.Ctx) error {
    var record CardRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateCardRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create card record",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(record)
}

// GetCards retrieves all cards
func (h *Handler) GetCards(c *fiber.Ctx) error {
    var records []CardRecord
    err := h.service.db.Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve card records",
        })
    }

    return c.JSON(records)
}

// GetCard retrieves a card by ID
func (h *Handler) GetCard(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    record, err := h.service.GetCardRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Card record not found",
        })
    }

    return c.JSON(record)
}

// UpdateCard updates a card
func (h *Handler) UpdateCard(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    var record CardRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    record.ID = uint(id)

    if err := h.service.UpdateCardRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update card record",
        })
    }

    return c.JSON(record)
}

// ActivateCard activates a card
func (h *Handler) ActivateCard(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    // In a real implementation, this would call the card issuing provider to activate the card
    // For now, we'll just update the status in our database
    
    record, err := h.service.GetCardRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Card record not found",
        })
    }
    
    record.Status = "active"
    if err := h.service.UpdateCardRecord(record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update card record",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Card activated successfully",
    })
}

// DeactivateCard deactivates a card
func (h *Handler) DeactivateCard(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    // In a real implementation, this would call the card issuing provider to deactivate the card
    // For now, we'll just update the status in our database
    
    record, err := h.service.GetCardRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Card record not found",
        })
    }
    
    record.Status = "inactive"
    if err := h.service.UpdateCardRecord(record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update card record",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Card deactivated successfully",
    })
}

// CancelCard cancels a card
func (h *Handler) CancelCard(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    // In a real implementation, this would call the card issuing provider to cancel the card
    // For now, we'll just update the status in our database
    
    record, err := h.service.GetCardRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Card record not found",
        })
    }
    
    record.Status = "cancelled"
    if err := h.service.UpdateCardRecord(record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update card record",
        })
    }

    return c.JSON(fiber.Map{
        "message": "Card cancelled successfully",
    })
}

// GetCardTransactions retrieves card transactions
func (h *Handler) GetCardTransactions(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid card ID",
        })
    }

    var records []TransactionRecord
    err = h.service.db.Where("card_id = ?", id).Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve transaction records",
        })
    }

    return c.JSON(records)
}

// CreateCardholder creates a new cardholder
func (h *Handler) CreateCardholder(c *fiber.Ctx) error {
    var record CardholderRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateCardholderRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create cardholder record",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(record)
}

// GetCardholders retrieves all cardholders
func (h *Handler) GetCardholders(c *fiber.Ctx) error {
    var records []CardholderRecord
    err := h.service.db.Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve cardholder records",
        })
    }

    return c.JSON(records)
}

// GetCardholder retrieves a cardholder by ID
func (h *Handler) GetCardholder(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid cardholder ID",
        })
    }

    record, err := h.service.GetCardholderRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Cardholder record not found",
        })
    }

    return c.JSON(record)
}

// UpdateCardholder updates a cardholder
func (h *Handler) UpdateCardholder(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid cardholder ID",
        })
    }

    var record CardholderRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    record.ID = uint(id)

    if err := h.service.UpdateCardholderRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update cardholder record",
        })
    }

    return c.JSON(record)
}