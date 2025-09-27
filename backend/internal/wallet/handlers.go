package wallet

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// Handler handles wallet HTTP requests
type Handler struct {
    service *Service
}

// NewHandler creates a new wallet handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{
        service: NewService(db),
    }
}

// SetupRoutes sets up the wallet routes
func SetupRoutes(router fiber.Router, db *gorm.DB) {
    handler := NewHandler(db)
    
    wallets := router.Group("/wallets")
    {
        wallets.Post("/", handler.CreateWallet)
        wallets.Get("/", handler.GetWallets)
        wallets.Get("/:id", handler.GetWallet)
        wallets.Put("/:id", handler.UpdateWallet)
        wallets.Delete("/:id", handler.DeleteWallet)
        
        // Transactions
        wallets.Get("/:id/transactions", handler.GetTransactions)
        wallets.Post("/:id/transactions", handler.CreateTransaction)
        
        // Custodial wallets
        wallets.Post("/custodial", handler.CreateCustodialWallet)
        wallets.Get("/custodial", handler.GetCustodialWallets)
    }
}

// CreateWallet creates a new wallet
func (h *Handler) CreateWallet(c *fiber.Ctx) error {
    var wallet Wallet
    if err := c.BodyParser(&wallet); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateWallet(&wallet); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create wallet",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(wallet)
}

// GetWallets retrieves all wallets for a user
func (h *Handler) GetWallets(c *fiber.Ctx) error {
    userID, err := strconv.Atoi(c.Query("user_id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    wallets, err := h.service.GetWalletsByUserID(uint(userID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve wallets",
        })
    }

    return c.JSON(wallets)
}

// GetWallet retrieves a wallet by ID
func (h *Handler) GetWallet(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid wallet ID",
        })
    }

    wallet, err := h.service.GetWalletByID(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Wallet not found",
        })
    }

    return c.JSON(wallet)
}

// UpdateWallet updates a wallet
func (h *Handler) UpdateWallet(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid wallet ID",
        })
    }

    var wallet Wallet
    if err := c.BodyParser(&wallet); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Ensure the ID matches
    wallet.ID = uint(id)

    if err := h.service.UpdateWallet(&wallet); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot update wallet",
        })
    }

    return c.JSON(wallet)
}

// DeleteWallet deletes a wallet
func (h *Handler) DeleteWallet(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid wallet ID",
        })
    }

    if err := h.service.DeleteWallet(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot delete wallet",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}

// CreateTransaction creates a new transaction
func (h *Handler) CreateTransaction(c *fiber.Ctx) error {
    walletID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid wallet ID",
        })
    }

    var tx Transaction
    if err := c.BodyParser(&tx); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    // Set the wallet ID
    tx.WalletID = uint(walletID)

    if err := h.service.CreateTransaction(&tx); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create transaction",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(tx)
}

// GetTransactions retrieves all transactions for a wallet
func (h *Handler) GetTransactions(c *fiber.Ctx) error {
    walletID, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid wallet ID",
        })
    }

    transactions, err := h.service.GetTransactionsByWalletID(uint(walletID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve transactions",
        })
    }

    return c.JSON(transactions)
}

// CreateCustodialWallet creates a new custodial wallet
func (h *Handler) CreateCustodialWallet(c *fiber.Ctx) error {
    var wallet CustodialWallet
    if err := c.BodyParser(&wallet); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    if err := h.service.CreateCustodialWallet(&wallet); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot create custodial wallet",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(wallet)
}

// GetCustodialWallets retrieves all custodial wallets for a user
func (h *Handler) GetCustodialWallets(c *fiber.Ctx) error {
    userID, err := strconv.Atoi(c.Query("user_id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid user ID",
        })
    }

    wallets, err := h.service.GetCustodialWalletsByUserID(uint(userID))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot retrieve custodial wallets",
        })
    }

    return c.JSON(wallets)
}