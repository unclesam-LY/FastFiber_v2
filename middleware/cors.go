package middleware

import (
	"FastFiber_v2/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func CORS() fiber.Handler {
	allowOrigins := strings.Join(global.Config.System.AllowOrigins, ",")

	return func(c *fiber.Ctx) error {
		cors.New(cors.Config{
			AllowOrigins:     allowOrigins,
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
			AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
			AllowCredentials: true,
			ExposeHeaders:    "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers",
			MaxAge:           86400,
		})
		return nil
	}
}
