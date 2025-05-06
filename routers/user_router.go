package routers

import (
	"FastFiber_v2/api"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(g fiber.Router) {
	app := api.App.UserApi
	g.Post("register", app.Register)
	g.Post("login", app.Login)
}
