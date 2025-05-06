package core

import (
	"FastFiber_v2/global"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() (client *redis.Client) {
	cfg := global.Config.Redis
	if cfg.Addr == "" {
		zap.L().Warn("未配置redis连接")
		return
	}

	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("redis连接失败", zap.Error(err))
		return
	}
	zap.L().Info("连接成功redis")
	return
}
