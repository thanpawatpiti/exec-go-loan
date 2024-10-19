package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers"
)

func AuthRoutes(api fiber.Router, handler handlers.InitHandler) {
	auth := api.Group("/auth")

	auth.Post("/login", handler.Login)
	auth.Post("/refresh", handler.Refresh)
}
