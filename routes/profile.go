package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers"

	auth "github.com/thanpawatpiti/exec-go-loan/routes/middlewares"
)

func ProfileRoute(api fiber.Router, handler handlers.InitHandler) {
	auth := api.Group("/profile").Use(auth.Protected())

	auth.Get("/", handler.GetProfile)
}
