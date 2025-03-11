package log

import "fmt"

type Level int8

const (
	TraceLevel Level = iota - 2
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// String returns a lower-case ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// LevelForGorm 转换成gorm日志级别
func (l Level) LevelForGorm() int {
	switch l {
	case FatalLevel, ErrorLevel:
		return 2
	case WarnLevel:
		return 3
	case InfoLevel, DebugLevel, TraceLevel:
		return 4
	default:
		return 1
	}
}

func GetLevel(levelStr string) (Level, error) {
	switch levelStr {
	case TraceLevel.String():
		return TraceLevel, nil
	case DebugLevel.String():
		return DebugLevel, nil
	case InfoLevel.String():
		return InfoLevel, nil
	case WarnLevel.String():
		return WarnLevel, nil
	case ErrorLevel.String():
		return ErrorLevel, nil
	case FatalLevel.String():
		return FatalLevel, nil
	}
	return InfoLevel, fmt.Errorf("unknown Level String: '%s', defaulting to InfoLevel", levelStr)
}

// Opts 获取logger的全局配置
func Opts() *Options {
	return std.opts
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
