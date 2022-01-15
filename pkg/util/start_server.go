package util

import (
	"fmt"
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/cenkkoroglu/oz-fiber/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnectionsClosed := make(chan struct{})

	config := config.Config

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			logger.Error("Server is not shutting down!", zap.Error(err))
		}

		close(idleConnectionsClosed)
	}()

	logger.Info(fmt.Sprintf("Server is started :%d", config.Port))

	if err := a.Listen(fmt.Sprintf(":%d", config.Port)); err != nil {
		logger.Error("Server is not running!", zap.Error(err))
	}

	<-idleConnectionsClosed
}
