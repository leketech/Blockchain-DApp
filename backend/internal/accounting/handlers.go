package accounting

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles accounting HTTP requests
type Handler struct {
    service *LedgerService
}

// NewHandler creates a new accounting handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewLedgerService(db),
    }
}

// SetupRoutes sets up the accounting routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    accounting := router.Group("/accounting")
    {
        // Accounts
        accounting.Post("/accounts", handler.CreateAccount)
        accounting.Get("/accounts", handler.GetAccounts)
        accounting.Get("/accounts/:id", handler.GetAccount)
        accounting.Put("/accounts/:id", handler.UpdateAccount)
        accounting.Post("/accounts/:id/freeze", handler.FreezeAccount)
        accounting.Post("/accounts/:id/unfreeze", handler.UnfreezeAccount)
        
        // Transactions
        accounting.Post("/transactions", handler.CreateTransaction)
        accounting.Get("/transactions", handler.GetTransactions)
        accounting.Get("/transactions/:id", handler.GetTransaction)
        accounting.Put("/transactions/:id/status", handler.UpdateTransactionStatus)
        
        // Journal entries
        accounting.Post("/journal-entries", handler.CreateJournalEntry)
        accounting.Get("/journal-entries", handler.GetJournalEntries)
        
        // Balances
        accounting.Get("/accounts/:id/balance", handler.GetAccountBalance)
    }
}

// CreateAccount creates a new account
func (h *Handler) CreateAccount(c *fiber.Ctx) error {
    var account Account
    if err := c.BodyParser(&account); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdAccount, err := h.service.CreateAccount(c.Context(), account)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create account",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdAccount)
}

// GetAccounts retrieves all accounts
func (h *Handler) GetAccounts(c *fiber.Ctx) error {
    // TODO: Implement pagination
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// GetAccount retrieves an account by ID
func (h *Handler) GetAccount(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid account ID",
        })
    }

    account, err := h.service.GetAccount(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Account not found",
        })
    }

    return c.JSON(account)
}

// UpdateAccount updates an account
func (h *Handler) UpdateAccount(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid account ID",
        })
    }

    var account Account
    if err := c.BodyParser(&account); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    account.ID = uint(id)

    // TODO: Implement update logic
    return c.JSON(account)
}

// FreezeAccount freezes an account
func (h *Handler) FreezeAccount(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid account ID",
        })
    }

    if err := h.service.FreezeAccount(c.Context(), uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot freeze account",
        })
    }

    return c.JSON(fiber.Map{"message": "Account frozen successfully"})
}

// UnfreezeAccount unfreezes an account
func (h *Handler) UnfreezeAccount(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid account ID",
        })
    }

    if err := h.service.UnfreezeAccount(c.Context(), uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot unfreeze account",
        })
    }

    return c.JSON(fiber.Map{"message": "Account unfrozen successfully"})
}

// CreateTransaction creates a new transaction
func (h *Handler) CreateTransaction(c *fiber.Ctx) error {
    var transaction Transaction
    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdTransaction, err := h.service.CreateTransaction(c.Context(), transaction)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create transaction",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdTransaction)
}

// GetTransactions retrieves all transactions
func (h *Handler) GetTransactions(c *fiber.Ctx) error {
    // TODO: Implement pagination and filtering
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// GetTransaction retrieves a transaction by ID
func (h *Handler) GetTransaction(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid transaction ID",
        })
    }

    transaction, err := h.service.GetTransaction(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Transaction not found",
        })
    }

    return c.JSON(transaction)
}

// UpdateTransactionStatus updates a transaction's status
func (h *Handler) UpdateTransactionStatus(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid transaction ID",
        })
    }

    var request struct {
        Status TransactionStatus `json:"status"`
    }
    
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.UpdateTransactionStatus(c.Context(), uint(id), request.Status); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update transaction status",
        })
    }

    return c.JSON(fiber.Map{"message": "Transaction status updated successfully"})
}

// CreateJournalEntry creates a new journal entry
func (h *Handler) CreateJournalEntry(c *fiber.Ctx) error {
    var entry JournalEntry
    if err := c.BodyParser(&entry); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    createdEntry, err := h.service.CreateJournalEntry(c.Context(), entry)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create journal entry",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(createdEntry)
}

// GetJournalEntries retrieves all journal entries
func (h *Handler) GetJournalEntries(c *fiber.Ctx) error {
    // TODO: Implement filtering by transaction ID, account ID, etc.
    return c.JSON(fiber.Map{"message": "Not implemented"})
}

// GetAccountBalance retrieves the balance for an account
func (h *Handler) GetAccountBalance(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid account ID",
        })
    }

    balance, err := h.service.GetAccountBalance(c.Context(), uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Account not found",
        })
    }

    return c.JSON(fiber.Map{"balance": balance})
}