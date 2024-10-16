package routes

import (
	"github.com/thanpawatpiti/exec-go-loan/pkg"
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Route   *fiber.App
	Handler handlers.InitHandler
}

func NewRoute(cmdInterface pkg.CmdInterface, handler handlers.InitHandler) *Route {
	r := fiber.New()
	r.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	return &Route{
		Route:   r,
		Handler: handler,
	}
}

func (r *Route) Load() {
	api := r.Route.Group("/api/v1")

	AuthRoutes(api, r.Handler)
	HealthRoutes(api, r.Handler)
}

func (r *Route) Start() {
	if err := r.Route.Listen(":3000"); err != nil {
		panic(err)
	}
}
