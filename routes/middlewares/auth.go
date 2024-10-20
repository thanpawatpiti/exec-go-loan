package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	helpers "github.com/thanpawatpiti/exec-go-loan/pkg/helpers"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// แยก token ออกมาจาก "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}
		tokenStr := tokenParts[1]

		claims, err := helpers.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// ตรวจสอบ expiration
		if claims["exp"] == nil || int64(claims["exp"].(float64)) < time.Now().Unix() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token is expired"})
		}

		// ตรวจสอบ user_id
		if claims["user_id"] == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// ตั้งค่าข้อมูล user_id ใน context
		c.Locals("user_id", uint(claims["user_id"].(float64)))

		return c.Next()
	}
}
