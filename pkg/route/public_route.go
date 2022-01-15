package route

import (
	"github.com/cenkkoroglu/oz-fiber/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoute(app *fiber.App) {
	appHandler := handlers.NewAppHandler()

	app.Get("/api/app/health", appHandler.Health)
}
