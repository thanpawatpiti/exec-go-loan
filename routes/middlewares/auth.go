package middlewares

import (
	"github.com/gofiber/fiber/v2"

	helpers "github.com/thanpawatpiti/exec-go-loan/pkg/helpers"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		claims, err := helpers.ValidateToken(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Check if token is expired
		if claims["exp"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is expired"})
		}

		// Check if user is not found
		if claims["user_id"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Set user information to context for future use
		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}
