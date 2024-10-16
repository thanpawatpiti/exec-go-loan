package routes

import "github.com/gofiber/fiber/v2"

func HealthRoutes(app *Route) {
	app.Route.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
