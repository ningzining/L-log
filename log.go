package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

type Logger struct {
	log  *zap.Logger
	opts *Options
}

func New(opts ...Option) *Logger {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	logger := newLogger(o)

	return &Logger{log: logger, opts: o}
}

func (l *Logger) Options() *Options {
	return l.opts
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
	l.log.Log(zapcore.Level(lvl), msg, fields...)
}

func (l *Logger) Sync() {
	_ = l.log.Sync()
}

var std = New()

func Default() *Logger {
	return std
}

func ReplaceDefault(l *Logger) {
	std = l
}

func Sync() {
	std.Sync()
}
