package core

import (
	"FastFiber_v2/config"
	"FastFiber_v2/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strings"
)

func InitZapLogger() *zap.Logger {
	logConfig := global.Config.Log

	// 创建三个日志核心（info/warn共用、error单独、控制台）
	cores := []zapcore.Core{
		createCore("info_warn", logConfig.Filename, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel

		}), logConfig),
		createCore("error", strings.Replace(logConfig.Filename, ".log", "_error.log", 1), zap.ErrorLevel, logConfig),
		createConsoleCore(),
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())

}

// 创建控制台核心
func createCore(logType string, filename string, enab zapcore.LevelEnabler, cfg config.Log) zapcore.Core {
	return zapcore.NewCore(
		getEncoder(),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups, // 修正顺序
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}),
		enab,
	)
}

func createConsoleCore() zapcore.Core {
	level := zap.DebugLevel
	if global.Config.Log.Env == "release" {
		level = zap.InfoLevel // 生产环境控制台只输出info及以上
	}
	return zapcore.NewCore(
		getEncoder(),
		zapcore.AddSync(os.Stdout),
		level,
	)
}

// 获取统一编码器
func getEncoder() zapcore.Encoder {
	prefix := global.Config.Log.Prefix
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		FunctionKey:    zapcore.OmitKey, // 隐藏函数名
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customEncodeLevel,                                                              // 自定义级别格式
		EncodeTime:     zapcore.TimeEncoderOfLayout(fmt.Sprintf("[%s] [2006-01-02 15:04:05]", prefix)), // 时间格式
		EncodeCaller:   customCallerEncoder,                                                            // 自定义调用者格式
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	if global.Config.Log.Env == "release" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 自定义调用者格式（文件:行号 函数名）
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s:%d %s",
		path.Base(caller.File), // 文件名
		caller.Line,            // 行号
		caller.Function,        // 函数名
	))
}

// 自定义日志级别显示
func customEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = "\x1b[36m" // 青色
	case zapcore.InfoLevel:
		color = "\x1b[34m" // 绿色
	case zapcore.WarnLevel:
		color = "\x1b[33m" // 黄色
	case zapcore.ErrorLevel:
		color = "\x1b[31m" // 红色
	default:
		color = "\x1b[0m" // 默认色
	}

	levelStr := level.CapitalString()
	if levelStr == "" {
		levelStr = "UNKNOWN"
	}

	enc.AppendString(fmt.Sprintf("%s[%s]\x1b[0m", color, levelStr))
}
