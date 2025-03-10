package log

type Option func(*options)

type options struct {
	Level      Level  // 日志级别
	path       string // 日志文件路径
	stdout     string // 输出类型，default: 终端输出，file:文本输出
	maxSize    int    // 日志文件最大大小（MB）
	maxAge     int    // 日志文件最大保存天数
	maxBackups int    // 日志文件最大保存个数
	compress   bool   // 是否压缩日志文件
	format     string // 日志输出格式，default: 终端输出，json: json格式输出
}

func newDefault() *options {
	return &options{
		Level:      DebugLevel,
		path:       "temp/logs",
		stdout:     "default",
		maxSize:    10,
		maxAge:     7,
		maxBackups: 10,
		compress:   false,
		format:     "default",
	}
}

func WithLevel(s Level) Option {
	return func(o *options) {
		o.Level = s
	}
}

func WithPath(s string) Option {
	return func(o *options) {
		o.path = s
	}
}

func WithStdout(s string) Option {
	return func(o *options) {
		o.stdout = s
	}
}

func WithMaxSize(n int) Option {
	return func(o *options) {
		o.maxSize = n
	}
}

func WithMaxAge(n int) Option {
	return func(o *options) {
		o.maxAge = n
	}
}

func WithMaxBackups(n int) Option {
	return func(o *options) {
		o.maxBackups = n
	}
}

func WithCompress(compress bool) Option {
	return func(o *options) {
		o.compress = compress
	}
}
