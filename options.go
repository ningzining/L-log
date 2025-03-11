package log

type Option func(*Options)

const (
	stdout  = "default" // 终端输出
	fileOut = "file"    // 文件输出
)

const (
	defaultFormat = "default" // 终端输出
	jsonFormat    = "json"    // json格式输出
)

const (
	defaultPath = "temp/logs"
)

type Options struct {
	Level      Level  // 日志级别
	Path       string // 日志文件路径
	Out        string // 输出类型，default: 终端输出，file:文本输出
	MaxSize    int    // 日志文件最大大小（MB）
	MaxAge     int    // 日志文件最大保存天数
	MaxBackups int    // 日志文件最大保存个数
	Compress   bool   // 是否压缩日志文件
	Format     string // 日志输出格式，default: 终端输出，json: json格式输出
}

func defaultOptions() *Options {
	return &Options{
		Level:      DebugLevel,
		Path:       defaultPath,
		Out:        stdout,
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   false,
		Format:     defaultFormat,
	}
}

func WithLevel(s Level) Option {
	return func(o *Options) {
		o.Level = s
	}
}

func WithPath(s string) Option {
	return func(o *Options) {
		o.Path = s
	}
}

func WithOut(s string) Option {
	return func(o *Options) {
		o.Out = s
	}
}

func WithMaxSize(n int) Option {
	return func(o *Options) {
		o.MaxSize = n
	}
}

func WithMaxAge(n int) Option {
	return func(o *Options) {
		o.MaxAge = n
	}
}

func WithMaxBackups(n int) Option {
	return func(o *Options) {
		o.MaxBackups = n
	}
}

func WithCompress(compress bool) Option {
	return func(o *Options) {
		o.Compress = compress
	}
}
