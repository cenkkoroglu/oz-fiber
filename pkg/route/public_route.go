package route

import (
	"github.com/cenkkoroglu/oz-fiber/app/handlers"
	"github.com/cenkkoroglu/oz-fiber/app/services"
	"github.com/cenkkoroglu/oz-fiber/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func PublicRoute(app *fiber.App) {
	// App routes
	appHandler := handlers.NewAppHandler()

	app.Get("/api/app/health", appHandler.Health)

	// User routes
	userService := services.NewUserService(database.DBConn)
	userHandler := handlers.NewUserHandler(userService)

	app.Post("/api/user/login", userHandler.Login)
	app.Post("/api/user/register", userHandler.Register)
}
