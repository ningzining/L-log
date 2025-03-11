package log

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path/filepath"
	"time"
)

func NewLumberjackWriter(opt *Options) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath.Join(opt.Path, fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))),
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
		LocalTime:  true,
		Compress:   opt.Compress,
	}
}
