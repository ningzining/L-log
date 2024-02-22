package log

import (
	"fmt"

	"go.uber.org/zap"
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

type Logger struct {
	logs []*zap.Logger
}

func New(cnf Config) *Logger {
	cnf.complete()
	loggers := cnf.newZapLog()

	return &Logger{logs: loggers}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.output(DebugLevel, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.output(InfoLevel, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.output(WarnLevel, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.output(ErrorLevel, msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.output(PanicLevel, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.output(FatalLevel, msg, fields...)
}

func (l *Logger) Debugf(msg string, a ...any) {
	l.output(DebugLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) Infof(msg string, a ...any) {
	l.output(InfoLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) Warnf(msg string, a ...any) {
	l.output(WarnLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) Errorf(msg string, a ...any) {
	l.output(ErrorLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) Panicf(msg string, a ...any) {
	l.output(PanicLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) Fatalf(msg string, a ...any) {
	l.output(FatalLevel, fmt.Sprintf(msg, a...))
}

func (l *Logger) output(lvl Level, msg string, fields ...Field) {
	for _, log := range l.logs {
		log.Log(lvl, msg, fields...)
	}
}

func (l *Logger) Sync() {
	for _, log := range l.logs {
		_ = log.Sync()
	}
}

var std = New(Config{})

func Default() *Logger {
	return std
}

func ReplaceDefault(l *Logger) {
	std = l
}

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

func Sync() {
	std.Sync()
}
