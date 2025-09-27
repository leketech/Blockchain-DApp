package payments

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles payment HTTP requests
type Handler struct {
    service *Service
}

// NewHandler creates a new payment handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewService(db),
    }
}

// SetupRoutes sets up the payment routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    payments := router.Group("/payments")
    {
        payments.Post("/", handler.CreatePayment)
        payments.Get("/", handler.GetPayments)
        payments.Get("/:id", handler.GetPayment)
        payments.Put("/:id", handler.UpdatePayment)
        
        // Customers
        payments.Post("/customers", handler.CreateCustomer)
        payments.Get("/customers", handler.GetCustomers)
        payments.Get("/customers/:id", handler.GetCustomer)
        payments.Put("/customers/:id", handler.UpdateCustomer)
        
        // Payment methods
        payments.Post("/methods", handler.CreatePaymentMethod)
        payments.Get("/methods", handler.GetPaymentMethods)
        payments.Get("/methods/:id", handler.GetPaymentMethod)
        payments.Put("/methods/:id", handler.UpdatePaymentMethod)
    }
}

// CreatePayment creates a new payment
func (h *Handler) CreatePayment(c *fiber.Ctx) error {
    var record PaymentRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreatePaymentRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create payment record",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(record)
}

// GetPayments retrieves all payments
func (h *Handler) GetPayments(c *fiber.Ctx) error {
    var records []PaymentRecord
    err := h.service.db.Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve payment records",
        })
    }

    return c.JSON(records)
}

// GetPayment retrieves a payment by ID
func (h *Handler) GetPayment(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid payment ID",
        })
    }

    record, err := h.service.GetPaymentRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Payment record not found",
        })
    }

    return c.JSON(record)
}

// UpdatePayment updates a payment
func (h *Handler) UpdatePayment(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid payment ID",
        })
    }

    var record PaymentRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    record.ID = uint(id)

    if err := h.service.UpdatePaymentRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update payment record",
        })
    }

    return c.JSON(record)
}

// CreateCustomer creates a new customer
func (h *Handler) CreateCustomer(c *fiber.Ctx) error {
    var record CustomerRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateCustomerRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create customer record",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(record)
}

// GetCustomers retrieves all customers
func (h *Handler) GetCustomers(c *fiber.Ctx) error {
    var records []CustomerRecord
    err := h.service.db.Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve customer records",
        })
    }

    return c.JSON(records)
}

// GetCustomer retrieves a customer by ID
func (h *Handler) GetCustomer(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid customer ID",
        })
    }

    record, err := h.service.GetCustomerRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Customer record not found",
        })
    }

    return c.JSON(record)
}

// UpdateCustomer updates a customer
func (h *Handler) UpdateCustomer(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid customer ID",
        })
    }

    var record CustomerRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    record.ID = uint(id)

    if err := h.service.UpdateCustomerRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update customer record",
        })
    }

    return c.JSON(record)
}

// CreatePaymentMethod creates a new payment method
func (h *Handler) CreatePaymentMethod(c *fiber.Ctx) error {
    var record PaymentMethodRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreatePaymentMethodRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create payment method record",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(record)
}

// GetPaymentMethods retrieves all payment methods
func (h *Handler) GetPaymentMethods(c *fiber.Ctx) error {
    var records []PaymentMethodRecord
    err := h.service.db.Find(&records).Error
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve payment method records",
        })
    }

    return c.JSON(records)
}

// GetPaymentMethod retrieves a payment method by ID
func (h *Handler) GetPaymentMethod(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid payment method ID",
        })
    }

    record, err := h.service.GetPaymentMethodRecordByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Payment method record not found",
        })
    }

    return c.JSON(record)
}

// UpdatePaymentMethod updates a payment method
func (h *Handler) UpdatePaymentMethod(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid payment method ID",
        })
    }

    var record PaymentMethodRecord
    if err := c.BodyParser(&record); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    record.ID = uint(id)

    if err := h.service.UpdatePaymentMethodRecord(&record); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update payment method record",
        })
    }

    return c.JSON(record)
}