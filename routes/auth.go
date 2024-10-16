package routes

func AuthRoutes(app *Route) {
	auth := app.Route.Group("/auth")

	auth.Post("/login", app.Handler.Login)
}
