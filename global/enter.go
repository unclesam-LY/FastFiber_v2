package global

import (
	"FastFiber_v2/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Version = "1.0.1"

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *zap.Logger
)
