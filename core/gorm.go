package core

import (
	"FastFiber_v2/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

func InitGorm() (db *gorm.DB) {
	cfg := global.Config.DB
	var dialector = cfg.Dsn()
	if dialector == nil {
		return
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		zap.L().Fatal("数据库连接失败", zap.Error(err))
		return
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Fatal("获取数据库连接失败", zap.Error(err))
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		zap.L().Fatal("数据库连接失败", zap.Error(err))
		return
	}
	// 设置连接池
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	zap.L().Info("数据库连接成功")
	return
}
