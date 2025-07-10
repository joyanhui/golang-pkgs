package zapLogger

import (
	"fmt"
	"os"

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
