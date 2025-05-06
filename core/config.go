package core

import (
	"FastFiber_v2/config"
	"FastFiber_v2/flags"
	"FastFiber_v2/global"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		global.Log.Error("配置文件读取错误", zap.Error(err))
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		global.Log.Error("配置文件格式错误", zap.Error(err))
		return
	}
	log.Println("config yamlFile load Init success")
	return
}

func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error("配置文件转换错误", zap.String("path", flags.Options.File),
			zap.Error(err))
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		global.Log.Error("配置文件写入错误", zap.String("path", flags.Options.File),
			zap.Error(err))
		return
	}
	global.Log.Info("配置文件保存成功",
		zap.String("path", flags.Options.File))
}
