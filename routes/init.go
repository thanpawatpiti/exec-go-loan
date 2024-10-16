package routes

import (
	"github.com/thanpawatpiti/exec-go-loan/pkg"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Route        *fiber.App
	CmdInterface pkg.CmdInterface
	Model        pkg.ModelInterface
}

func NewRoute(cmdInterface pkg.CmdInterface, model pkg.ModelInterface) *Route {
	return &Route{
		Route:        fiber.New(),
		CmdInterface: cmdInterface,
		Model:        model,
	}
}

func (r *Route) Load() {
	AuthRoutes(r)
	HealthRoutes(r)
}

func (r *Route) Start() {
	if err := r.Route.Listen(":3000"); err != nil {
		panic(err)
	}
}
