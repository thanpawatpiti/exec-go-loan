package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers"
)

func HealthRoutes(api fiber.Router, handler handlers.InitHandler) {
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
