package routers

import (
	"FastFiber_v2/api"
	"FastFiber_v2/api/user_api"
	"FastFiber_v2/middleware"
	"FastFiber_v2/models"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(g fiber.Router) {
	app := api.App.UserApi
	g.Post("register", app.Register)
	g.Post("login",
		middleware.BindJsonMiddleware[user_api.Auth],
		app.Login)
	g.Get("users",
		middleware.BindQueryMiddleware[models.PageInfo],
		app.UserListView)
}
