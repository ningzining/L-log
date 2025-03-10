package log

import (
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

func Debug(msg string, fields ...Field) {
	std.Debug(msg, fields...)
}

func Debugf(msg string, a ...any) {
	std.Debugf(msg, a...)
}

func Info(msg string, fields ...Field) {
	std.Info(msg, fields...)
}

func Infof(msg string, a ...any) {
	std.Infof(msg, a...)
}

func Warn(msg string, fields ...Field) {
	std.Warn(msg, fields...)
}

func Warnf(msg string, a ...any) {
	std.Warnf(msg, a...)
}

func Error(msg string, fields ...Field) {
	std.Error(msg, fields...)
}

func Errorf(msg string, a ...any) {
	std.Errorf(msg, a...)
}

func Panic(msg string, fields ...Field) {
	std.Panic(msg, fields...)
}

func Panicf(msg string, a ...any) {
	std.Panicf(msg, a...)
}

func Fatal(msg string, fields ...Field) {
	std.Fatal(msg, fields...)
}

func Fatalf(msg string, a ...any) {
	std.Fatalf(msg, a...)
}
