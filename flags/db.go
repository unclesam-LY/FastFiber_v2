package flags

import (
	"FastFiber_v2/global"
	"FastFiber_v2/models"
	"go.uber.org/zap"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		zap.L().Error("表结构迁移失败", zap.Error(err))
		return
	}
	zap.L().Info("表结构迁移成功")
}
