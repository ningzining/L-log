package L_log

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level      Level  // 日志级别, 默认为info
	Filename   string // 日志文件路径,默认为空,只输出到控制台
	MaxSize    int    // 最大字节大小，默认10MB
	MaxAge     int    // 最大天数，默认7天
	MaxBackups int    // 最大备份文件数，默认10
}

func (c *Config) complete() {
	if c.MaxSize <= 0 {
		c.MaxSize = 1024 * 1024 * 10 // 10MB
	}
	if c.MaxAge <= 0 {
		c.MaxAge = 7
	}
	if c.MaxBackups <= 0 {
		c.MaxBackups = 10
	}
}

func (c *Config) newZapLog() []*zap.Logger {
	loggers := make([]*zap.Logger, 0)
	loggers = append(loggers, c.newConsoleLogger())
	if c.Filename != "" {
		loggers = append(loggers, c.newJsonLogger())
	}

	return loggers
}

func (c *Config) newConsoleLogger() *zap.Logger {
	return zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(buildZapProductionConfig().EncoderConfig),
				zapcore.AddSync(os.Stdout),
				buildLevelEnablerFunc(zap.NewAtomicLevelAt(c.Level).Level()),
			),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.AddStacktrace(buildErrorLevelEnablerFunc()),
	)
}

func (c *Config) newJsonLogger() *zap.Logger {
	return zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(buildZapProductionConfig().EncoderConfig),
				zapcore.AddSync(c.newInfoLumberjack()),
				buildLevelEnablerFunc(zap.NewAtomicLevelAt(c.Level).Level()),
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(buildZapProductionConfig().EncoderConfig),
				zapcore.AddSync(c.newErrorLumberjack()),
				buildErrorLevelEnablerFunc(),
			),
		),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.AddStacktrace(buildErrorLevelEnablerFunc()),
	)
}

func (c *Config) newInfoLumberjack() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s-%s", c.Filename, "info"),
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		LocalTime:  true,
		Compress:   true,
	}
}

func (c *Config) newErrorLumberjack() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s-%s", c.Filename, "error"),
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		LocalTime:  true,
		Compress:   true,
	}
}
