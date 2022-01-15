package middleware

import (
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/cenkkoroglu/oz-fiber/pkg/util"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtMiddleware() fiber.Handler {
	config := config.GetConfig()
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.JwtSecret),
		SigningMethod: "HS256",
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return util.RestError(ctx, 401, []string{"Unauthorized"})
		},
	})
}
