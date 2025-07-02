package routers

import (
	"FastFiber_v2/global"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run() {
	r := fiber.New()

	// 添加中间件

	// 静态文件路由
	r.Static("/uploads", "./uploads")
	g := r.Group("api")

	// 注册路由
	UserRouter(g)
	CaptchaRouterGroup(g)

	// 获取监听地址
	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		global.Log.Sugar().Infof("后端服务运行在 %s", addr)
	}

	err := r.Listen(addr)
	if err != nil {
		zap.L().Error("服务启动错误:", zap.Error(err))
		return
	}
}
