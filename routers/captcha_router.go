package routers

import (
	"FastFiber_v2/api"
	"github.com/gofiber/fiber/v2"
)

func CaptchaRouterGroup(g fiber.Router) {
	app := api.App.CaptchaApi
	g.Get("captcha", app.GenerateView)
}
