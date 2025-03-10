package log

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path/filepath"
	"time"
)

func NewLumberjackWriter(opt *options) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath.Join(opt.path, fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))),
		MaxSize:    opt.maxSize,
		MaxAge:     opt.maxAge,
		MaxBackups: opt.maxBackups,
		LocalTime:  true,
		Compress:   opt.compress,
	}
}
