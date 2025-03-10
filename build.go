package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func newLogger(opt *options) *zap.Logger {
	var writer zapcore.WriteSyncer
	switch opt.stdout {
	case "file":
		writer = zapcore.AddSync(NewLumberjackWriter(opt))
	default:
		writer = os.Stdout
	}

	var encoder zapcore.Encoder
	switch opt.format {
	case "json":
		encoder = newJsonEncoder()
	default:
		encoder = newConsoleEncoder()
	}

	return zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				encoder,
				writer,
				buildLevelEnablerFunc(zap.NewAtomicLevelAt(opt.Level).Level()),
			),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.AddStacktrace(buildErrorLevelEnablerFunc()),
	)
}

func buildErrorLevelEnablerFunc() zap.LevelEnablerFunc {
	return buildLevelEnablerFunc(ErrorLevel)
}

func buildLevelEnablerFunc(lvl Level) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= lvl
	}
}
