package route

import (
	"github.com/cenkkoroglu/oz-fiber/app/handlers"
	"github.com/cenkkoroglu/oz-fiber/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoute(app *fiber.App) {
	appHandler := handlers.NewAppHandler()

	// JWT Middleware
	privateApp := app.Group("", middleware.JwtMiddleware())

	privateApp.Get("/api/app/health2", appHandler.Health2)
}
