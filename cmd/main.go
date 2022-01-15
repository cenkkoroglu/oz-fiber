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

	if databaseErr := database.Init(); databaseErr != nil {
		logger.Fatal("DB initialization failed!", zap.Error(databaseErr))
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(recover.New())
	app.Use(cors.New())

	if config.Environment != "production" {
		app.Get("/swagger/*", fiberSwagger.WrapHandler)
	}

	route.PublicRoute(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("")
	})

	util.StartServerWithGracefulShutdown(app)
}
