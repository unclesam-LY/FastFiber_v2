package user_api

import (
	"FastFiber_v2/global"
	"FastFiber_v2/middleware"
	"FastFiber_v2/models"
	"FastFiber_v2/service/common"
	"FastFiber_v2/utils/res"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (UserApi) UserListView(c *fiber.Ctx) error {

	var cr = middleware.GetBind[models.PageInfo](c)

	list, count, err := common.QueryList(models.UserModel{}, common.QueryOptions{
		PageInfo: cr,
		Likes:    []string{"username"},
		Debug:    false,
	})

	if err != nil {
		global.Log.Error("用户列表查询失败", zap.Error(err))
		return res.FailWithMsg("用户列表查询失败", c)

	}

	return res.OkWithList(list, count, c)
}
