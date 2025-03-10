package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func newConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(buildZapProductionConfig().EncoderConfig)
}

func newJsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(buildZapProductionConfig().EncoderConfig)
}

func buildZapProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.Format(time.DateTime))
	}
	return config
}
