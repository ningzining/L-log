package L_log

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("hello world", zap.String("name", "cotton"))
}

func TestInfo(t *testing.T) {
	Info("hello world", zap.String("name", "cotton"))
}

func TestWarn(t *testing.T) {
	Warn("hello world", zap.String("name", "cotton"))
}

func TestError(t *testing.T) {
	Error("hello world", zap.String("name", "cotton"))
}

func TestPanic(t *testing.T) {
	Panic("hello world", zap.String("name", "cotton"))
}

func TestFatal(t *testing.T) {
	Fatal("hello world", zap.String("name", "cotton"))
}

func TestSetLevel(t *testing.T) {
	Info("hello world", zap.String("name", "cotton"))
	SetLevel(ErrorLevel)
	Info("hello world", zap.String("name", "cotton"))
}
