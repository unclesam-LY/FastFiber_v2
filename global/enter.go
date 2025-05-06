package global

import (
	"FastFiber_v2/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Version = "1.0.1"

var (
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
	Log    *zap.Logger
)
