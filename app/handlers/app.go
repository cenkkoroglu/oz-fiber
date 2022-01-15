package handlers

import (
	"github.com/cenkkoroglu/oz-fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type appHandler struct {
}

func (a *appHandler) Health(ctx *fiber.Ctx) error {
	return util.RestOk(ctx, 200, "Healthy!")
}

func (a *appHandler) Health2(ctx *fiber.Ctx) error {
	return util.RestOk(ctx, 200, "Healthy!")
}

type AppHandler interface {
	Health(ctx *fiber.Ctx) error
	Health2(ctx *fiber.Ctx) error
}

func NewAppHandler() AppHandler {
	return &appHandler{}
}
