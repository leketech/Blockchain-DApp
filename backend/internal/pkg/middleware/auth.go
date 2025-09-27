package middleware

import (
    "github.com/blockchain-dapp/backend/internal/pkg/config"
    "github.com/blockchain-dapp/backend/internal/pkg/logger"
    "github.com/blockchain-dapp/backend/internal/pkg/security"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
)

// AuthMiddleware handles JWT authentication
func AuthMiddleware(cfg *config.Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get the JWT token from the Authorization header
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            logger.Warn("Missing authorization header")
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Missing authorization header",
            })
        }

        // Check if the header starts with "Bearer "
        if len(authHeader) < 7 || utils.ToLower(authHeader[:7]) != "bearer " {
            logger.Warn("Invalid authorization header format")
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid authorization header format",
            })
        }

        // Extract the token
        token := authHeader[7:]

        // Initialize JWT service
        jwtService := security.NewJWT(security.DefaultJWTConfig())

        // Validate the token
        claims, err := jwtService.ValidateToken(token)
        if err != nil {
            logger.Warn("Invalid token: " + err.Error())
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid token",
            })
        }

        // Set user information in context
        c.Locals("user_id", claims.UserID)
        c.Locals("email", claims.Email)
        c.Locals("role", claims.Role)
        c.Locals("full_name", claims.FullName)

        // Continue to the next middleware/handler
        return c.Next()
    }
}