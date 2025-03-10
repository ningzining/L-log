package log

type Option func(*Options)

const (
	stdoutDefault = "default" // 终端输出
	stdoutFile    = "file"    // 文件输出
)

const (
	formatDefault = "default" // 终端输出
	formatJson    = "json"    // json格式输出
)

const (
	pathDefault = "temp/logs"
)

type Options struct {
	Level      Level  // 日志级别
	path       string // 日志文件路径
	stdout     string // 输出类型，default: 终端输出，file:文本输出
	maxSize    int    // 日志文件最大大小（MB）
	maxAge     int    // 日志文件最大保存天数
	maxBackups int    // 日志文件最大保存个数
	compress   bool   // 是否压缩日志文件
	format     string // 日志输出格式，default: 终端输出，json: json格式输出
}

func defaultOptions() *Options {
	return &Options{
		Level:      DebugLevel,
		path:       pathDefault,
		stdout:     stdoutDefault,
		maxSize:    10,
		maxAge:     7,
		maxBackups: 10,
		compress:   false,
		format:     formatDefault,
	}
}

func WithLevel(s Level) Option {
	return func(o *Options) {
		o.Level = s
	}
}

func WithPath(s string) Option {
	return func(o *Options) {
		o.path = s
	}
}

func WithStdout(s string) Option {
	return func(o *Options) {
		o.stdout = s
	}
}

func WithMaxSize(n int) Option {
	return func(o *Options) {
		o.maxSize = n
	}
}

func WithMaxAge(n int) Option {
	return func(o *Options) {
		o.maxAge = n
	}
}

func WithMaxBackups(n int) Option {
	return func(o *Options) {
		o.maxBackups = n
	}
}

func WithCompress(compress bool) Option {
	return func(o *Options) {
		o.compress = compress
	}
}
