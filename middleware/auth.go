package middleware

import (
	"FastFiber_v2/service/redis_ser"
	"FastFiber_v2/utils/jwts"
	"FastFiber_v2/utils/res"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("access_token")
		_, err := jwts.CheckAccessToken(accessToken)
		if err != nil {
			return res.FailWithMsg("认证失败", c)
		}
		if redis_ser.HashLogout(accessToken); err != nil {
			return res.FailWithMsg("当前登录已经注销", c)
		}
		return c.Next()
	}
}

func AuthAdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("access_token")
		claims, err := jwts.CheckAccessToken(accessToken)
		if err != nil {
			return res.FailWithMsg("认证失败", c)
		}
		if redis_ser.HashLogout(accessToken) {
			return res.FailWithMsg("当前登录已经注销", c)
		}
		if claims.RoleID != 1 {
			return res.FailWithMsg("认证失败", c)
		}
		c.Locals("claims", claims)
		return c.Next()
	}
}

func GetAuth(c *fiber.Ctx) (cl *jwts.AccessPayload) {
	cl = new(jwts.AccessPayload)
	_claims := c.Locals("claims")
	if _claims == nil {
		return nil
	}

	cl, ok := _claims.(*jwts.AccessPayload)
	if !ok {
		return nil
	}

	return cl
}
