package log

type Level int8

// LevelForGorm 转换成gorm日志级别
func (l Level) LevelForGorm() int {
	switch l {
	case FatalLevel, ErrorLevel:
		return 2
	case WarnLevel:
		return 3
	case InfoLevel, DebugLevel:
		return 4
	default:
		return 1
	}
}

const (
	DebugLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

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
