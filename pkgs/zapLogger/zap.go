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
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间编码格式
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // 级别编码格式，例如 INFO, DEBUG

		config := zap.NewProductionConfig()
		config.EncoderConfig = encoderCfg            // 应用编码器配置
		config.OutputPaths = []string{"stderr"}      // 日志输出路径，这里是标准错误输出
		config.Encoding = "console"                  // 输出格式：console 或 json
		config.ErrorOutputPaths = []string{"stderr"} // 错误日志输出路径
		logLevelStr := os.Getenv("LOG_LEVEL")        // 日志登基
		effectiveLogLevel := parseLogLevel(logLevelStr)
		config.Level = zap.NewAtomicLevelAt(effectiveLogLevel) // 根据解析到的级别设置日志级别
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
		fmt.Println("isDebug")
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
