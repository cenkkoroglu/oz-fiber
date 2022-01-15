package main

import (
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/cenkkoroglu/oz-fiber/pkg/database"
	"github.com/cenkkoroglu/oz-fiber/pkg/logger"
	"github.com/cenkkoroglu/oz-fiber/pkg/middleware"
	"github.com/cenkkoroglu/oz-fiber/pkg/route"
	"github.com/cenkkoroglu/oz-fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

// @title Go OZ-Fiber Boilerplate
// @version 1.0
// @description OZ-Fiber is an easy-to-use Go Fiber Boilerplate project.
func main() {
	logger, loggerErr := logger.Init()
	if loggerErr != nil {
		panic("Failed to initialize logger.")
	}
	defer logger.Sync()

	if configErr := config.Init(); configErr != nil {
		logger.Fatal("Config is not loaded!", zap.Error(configErr))
	}

	config := config.GetConfig()

	// Connect database

	if databaseErr := database.Init(); databaseErr != nil {
		logger.Fatal("DB initialization failed!", zap.Error(databaseErr))
	}

	// Make database migrations

	if migrateErr := database.Migrate(); migrateErr != nil {
		logger.Fatal("DB migration failed!", zap.Error(migrateErr))
	}

	// Init validator
	util.InitValidator()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Compression middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Error recovering middleware
	app.Use(recover.New())

	// CORS middleware
	app.Use(cors.New())

	// Swagger route
	if config.Environment != "production" {
		app.Get("/swagger/*", fiberSwagger.WrapHandler)
	}

	// Public routes
	route.PublicRoute(app)

	// Private routes
	route.PrivateRoute(app)

	// Not found route
	app.Use(func(ctx *fiber.Ctx) error {
		return util.RestError(ctx, 404, []string{"Not found"})
	})

	util.StartServerWithGracefulShutdown(app)
}
