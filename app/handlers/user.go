package handlers

import (
	"github.com/cenkkoroglu/oz-fiber/app/models/request_models"
	"github.com/cenkkoroglu/oz-fiber/app/services"
	"github.com/cenkkoroglu/oz-fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type userHandler struct {
	userService services.UserService
}

func (u *userHandler) Login(ctx *fiber.Ctx) error {
	var payload request_models.Login

	if err := ctx.BodyParser(&payload); err != nil {
		return util.RestError(ctx, 400, []string{"Bad request"})
	}

	if err := util.ValidateStruct(payload); len(err) > 0 {
		return util.RestError(ctx, 400, err)
	}

	// Get user by email if it returns error, return invalid credentials error
	if user, err := u.userService.GetByEmail(payload.Email); err != nil {
		return util.RestError(ctx, 401, []string{"Invalid user credentials"})
	} else {
		// Check password if it returns error, return invalid credentials error
		if checkPassword := user.CheckPassword(payload.Password); checkPassword != nil {
			return util.RestError(ctx, 401, []string{"Invalid user credentials"})
		} else {
			// Generate token if it returns error, return invalid credentials error
			token, tokenErr := u.userService.GenerateToken(user)
			if tokenErr != nil {
				return util.RestError(ctx, 401, []string{"Invalid user credentials"})
			}

			return util.RestOk(ctx, 200, token)
		}
	}
}

func (u *userHandler) Register(ctx *fiber.Ctx) error {
	var payload request_models.Register

	if err := ctx.BodyParser(&payload); err != nil {
		return util.RestError(ctx, 400, []string{"Bad request"})
	}

	if err := util.ValidateStruct(payload); len(err) > 0 {
		return util.RestError(ctx, 400, err)
	}

	if _, err := u.userService.Register(&payload); err != nil {
		return util.RestError(ctx, 500, []string{"Error while creating user"})
	} else {
		return util.RestOk(ctx, 200, true)
	}
}

func NewUserHandler(_userService services.UserService) UserHandler {
	return &userHandler{
		userService: _userService,
	}
}
