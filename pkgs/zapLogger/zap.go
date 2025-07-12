package zapLogger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLogger(isDevMode bool) (*zap.SugaredLogger, error) {
	var logger *zap.SugaredLogger
	{
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

		config := zap.NewProductionConfig()
		config.EncoderConfig = encoderCfg
		config.OutputPaths = []string{"stderr"}
		config.Encoding = "console" // 输出格式 console 或 json
		config.ErrorOutputPaths = []string{"stderr"}
		logLevelStr := os.Getenv("LOG_LEVEL") // 日志登基
		effectiveLogLevel := parseLogLevel(logLevelStr)
		config.Level = zap.NewAtomicLevelAt(effectiveLogLevel)
		if isDevMode {
			config.DisableCaller = false
		} else {
			config.DisableCaller = true
		}

		plainLogger, err := config.Build()
		if err != nil {
			fmt.Printf("cant init zapLogger: %v\n", err)
			os.Exit(1)
		}
		defer plainLogger.Sync()
		logger = plainLogger.Sugar()
	}
	return logger, nil
}

func parseLogLevel(levelStr string) zapcore.Level {
	switch strings.ToUpper(levelStr) { // 将输入字符串转换为大写进行匹配
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN", "WARNING": // 同时支持 "WARN" 和 "WARNING"
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "DPANIC":
		return zapcore.DPanicLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		// 如果环境变量未设置或解析失败，默认日志级别为 Info
		return zapcore.InfoLevel
	}
}
