package L_log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildZapProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.Format(time.DateTime))
	}
	return config
}

func buildErrorLevelEnablerFunc() zap.LevelEnablerFunc {
	return buildLevelEnablerFunc(ErrorLevel)
}

func buildLevelEnablerFunc(lvl Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= lvl
	}
}
