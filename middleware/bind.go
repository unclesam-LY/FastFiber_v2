package middleware

import (
	"FastFiber_v2/utils/res"
	"github.com/gofiber/fiber/v2"
)

func BindJsonMiddleware[T any](c *fiber.Ctx) error {
	var cr T
	err := c.BodyParser(&cr)
	if err != nil {
		return res.FailWithError(err, c)
	}
	c.Locals("request", cr)

	return c.Next()
}

func BindQueryMiddleware[T any](c *fiber.Ctx) error {
	var cr T
	err := c.QueryParser(&cr)
	if err != nil {
		return res.FailWithError(err, c)
	}
	c.Locals("request", cr)
	return c.Next()
}

func BindUriMiddleware[T any](c *fiber.Ctx) error {
	var cr T
	err := c.ParamsParser(&cr)
	if err != nil {
		return res.FailWithError(err, c)

	}
	c.Locals("request", cr)
	return c.Next()
}

func GetBind[T any](c *fiber.Ctx) (cr T) {
	return c.Locals("request").(T)
}
