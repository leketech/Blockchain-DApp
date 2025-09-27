package admin

import (
    "fmt"
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles admin HTTP requests
type Handler struct {
    service *AdminService
}

// NewHandler creates a new admin handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewAdminService(db),
    }
}

// SetupRoutes sets up the admin routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    admin := router.Group("/admin")
    {
        // Support tickets
        admin.Post("/tickets", handler.CreateSupportTicket)
        admin.Get("/tickets", handler.GetSupportTickets)
        admin.Get("/tickets/:id", handler.GetSupportTicket)
        admin.Put("/tickets/:id", handler.UpdateSupportTicket)
        
        // Support messages
        admin.Post("/tickets/:ticketId/messages", handler.AddSupportMessage)
        admin.Get("/tickets/:ticketId/messages", handler.GetSupportMessages)
        
        // Audit logs
        admin.Get("/audit-logs", handler.GetAuditLogs)
        
        // System metrics
        admin.Post("/metrics", handler.CreateSystemMetric)
        admin.Get("/metrics", handler.GetSystemMetrics)
    }
}

// CreateSupportTicket creates a new support ticket
func (h *Handler) CreateSupportTicket(c *fiber.Ctx) error {
    var ticket SupportTicket
    if err := c.BodyParser(&ticket); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdTicket, err := h.service.CreateSupportTicket(c.Context(), ticket)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create support ticket",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdTicket)
}

// GetSupportTickets retrieves all support tickets
func (h *Handler) GetSupportTickets(c *fiber.Ctx) error {
    // TODO: Implement pagination and filtering
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// GetSupportTicket retrieves a support ticket by ID
func (h *Handler) GetSupportTicket(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ticket ID",
        })
    }

    ticket, err := h.service.GetSupportTicket(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Support ticket not found",
        })
    }

    return c.JSON(ticket)
}

// UpdateSupportTicket updates a support ticket
func (h *Handler) UpdateSupportTicket(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ticket ID",
        })
    }

    var updates map[string]interface{}
    if err := c.BodyParser(&updates); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.UpdateSupportTicket(c.Context(), uint(id), updates); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update support ticket",
        })
    }

    return c.JSON(fiber.Map{"message": "Support ticket updated successfully"})
}

// AddSupportMessage adds a message to a support ticket
func (h *Handler) AddSupportMessage(c *fiber.Ctx) error {
    ticketId, err := strconv.Atoi(c.Params("ticketId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ticket ID",
        })
    }

    var message SupportMessage
    if err := c.BodyParser(&message); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    message.TicketID = uint(ticketId)

    createdMessage, err := h.service.AddSupportMessage(c.Context(), message)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot add support message",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdMessage)
}

// GetSupportMessages retrieves messages for a support ticket
func (h *Handler) GetSupportMessages(c *fiber.Ctx) error {
    ticketId, err := strconv.Atoi(c.Params("ticketId"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ticket ID",
        })
    }

    messages, err := h.service.GetSupportMessages(c.Context(), uint(ticketId))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve support messages",
        })
    }

    return c.JSON(messages)
}

// CreateAuditLog creates a new audit log entry
func (h *Handler) CreateAuditLog(c *fiber.Ctx) error {
    var log AuditLog
    if err := c.BodyParser(&log); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateAuditLog(c.Context(), log); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create audit log",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(log)
}

// GetAuditLogs retrieves audit logs
func (h *Handler) GetAuditLogs(c *fiber.Ctx) error {
    // TODO: Implement pagination and filtering
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// CreateSystemMetric creates a new system metric
func (h *Handler) CreateSystemMetric(c *fiber.Ctx) error {
    var metric SystemMetric
    if err := c.BodyParser(&metric); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateSystemMetric(c.Context(), metric); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create system metric",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(metric)
}

// GetSystemMetrics retrieves system metrics
func (h *Handler) GetSystemMetrics(c *fiber.Ctx) error {
    // TODO: Implement pagination
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// ValidateSupportTicket validates a support ticket
func ValidateSupportTicket(ticket SupportTicket) error {
    if ticket.UserID == 0 {
        return fmt.Errorf("user ID is required")
    }
    
    if ticket.Subject == "" {
        return fmt.Errorf("subject is required")
    }
    
    if ticket.Description == "" {
        return fmt.Errorf("description is required")
    }
    
    if ticket.Category == "" {
        return fmt.Errorf("category is required")
    }
    
    return nil
}

// ValidateSupportMessage validates a support message
func ValidateSupportMessage(message SupportMessage) error {
    if message.TicketID == 0 {
        return fmt.Errorf("ticket ID is required")
    }
    
    if message.UserID == 0 {
        return fmt.Errorf("user ID is required")
    }
    
    if message.Message == "" {
        return fmt.Errorf("message is required")
    }
    
    return nil
}