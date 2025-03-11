package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func newLogger(opt *Options) *zap.Logger {
	return zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				buildEncoder(opt),
				buildWriteSyncer(opt),
				buildLevelEnabler(opt.Level),
			),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.AddStacktrace(buildLevelEnabler(ErrorLevel)),
	)
}

func buildEncoder(opt *Options) zapcore.Encoder {
	switch opt.Format {
	case jsonFormat:
		return zapcore.NewJSONEncoder(buildZapProductionConfig().EncoderConfig)
	default:
		return zapcore.NewConsoleEncoder(buildZapProductionConfig().EncoderConfig)
	}
}

func buildWriteSyncer(opt *Options) zapcore.WriteSyncer {
	switch opt.Out {
	case fileOut:
		return zapcore.AddSync(NewLumberjackWriter(opt))
	default:
		return os.Stdout
	}
}

func buildZapProductionConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.Format(time.DateTime))
	}
	return config
}

func buildLevelEnabler(lvl Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= zapcore.Level(lvl)
	}
}
