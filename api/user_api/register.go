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

type RegisterReq struct {
	Username   string `json:"username" validate:"required" label:"用户名"`
	RePassword string `json:"re_password" validate:"required" label:"二次密码"`
	Password   string `json:"password" validate:"required" label:"密码"`
}

func (UserApi) Register(c *fiber.Ctx) error {

	var cr = middleware.GetBind[RegisterReq](c)

	if cr.Password != cr.RePassword {
		return res.FailWithMsg("两次密码输入不一致", c)
	}

	var user models.UserModel
	err := global.DB.Take(&user, "username = ?", cr.Username).Error

	if err == nil {
		return res.FailWithMsg("用户名已存在", c)
	}

	hashPassword, err := pwd.GenerateFromPassword(cr.Password)
	if err != nil {
		global.Log.Error("密码加密失败", zap.Error(err))
		return res.FailWithMsg("注册失败", c)
	}

	err = global.DB.Create(&models.UserModel{
		Username: cr.Username,
		Password: hashPassword,
		RoleID:   2,
	}).Error

	if err != nil {
		global.Log.Error("创建用户失败", zap.Error(err))
		return res.FailWithMsg("用户注册失败", c)
	}

	accessToken, _ := jwts.GenAccessToken(jwts.Claims{
		UserID: user.ID,
		RoleID: uint(user.RoleID),
	})

	refreshToken, _ := jwts.GenRefreshToken(jwts.RefreshClaims{UserID: user.ID})

	return res.OkWithData(Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, c)
}
