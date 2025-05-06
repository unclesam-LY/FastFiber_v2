package user_api

import (
	"FastFiber_v2/global"
	"FastFiber_v2/middleware"
	"FastFiber_v2/models"
	"FastFiber_v2/utils/jwts"
	"FastFiber_v2/utils/pwd"
	"FastFiber_v2/utils/res"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Auth struct {
	Username string `json:"username" validate:"required" label:"用户名"`
	Password string `json:"password" validate:"required" label:"密码"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (UserApi) Login(c *fiber.Ctx) error {

	var cr = middleware.GetBind[Auth](c)

	var user models.UserModel
	err := global.DB.Where("username = ?", cr.Username).First(&user).Error

	if err != nil {
		return res.FailWithMsg("用户名或密码错误", c)
	}

	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		return res.FailWithMsg("用户名密码错误", c)
	}

	accessToken, err := jwts.GenAccessToken(jwts.Claims{
		UserID: user.ID,
		RoleID: uint(user.RoleID),
	})
	if err != nil {
		global.Log.Error("生成token失败", zap.Error(err))
		return res.FailWithMsg("登录失败", c)

	}
	refreshToken, err := jwts.GenRefreshToken(jwts.RefreshClaims{
		UserID: user.ID,
	})
	if err != nil {
		global.Log.Error("生成refreshToken失败", zap.Error(err))
		return res.FailWithMsg("登录失败", c)

	}

	return res.OkWithData(Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, c)

}
