package middleware

import (
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/cenkkoroglu/oz-fiber/pkg/logger"
	"github.com/cenkkoroglu/oz-fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	logger.Error(
		"System error!",
		zap.Error(err),
		zap.String("url", ctx.Request().URI().String()),
	)

	config := config.GetConfig()

	if config.Environment == "production" {
		return util.RestError(ctx, 500, []string{"Error occurred, please try again later!"})
	} else {
		return util.RestError(ctx, 500, []string{err.Error()})
	}
}
