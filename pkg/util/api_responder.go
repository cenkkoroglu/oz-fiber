package util

import (
	"github.com/cenkkoroglu/oz-fiber/app/models"
	"github.com/gofiber/fiber/v2"
)

func RestOk(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(models.ApiResponse{
		Success: true,
		Data:    data,
	})
}

func RestError(ctx *fiber.Ctx, status int, errors []string) error {
	return ctx.Status(status).JSON(models.ApiResponse{
		Success: false,
		Errors:  errors,
	})
}
