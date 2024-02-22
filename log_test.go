package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	Debug("hello world", zap.String("name", "cotton"))
}

func TestDebugf(t *testing.T) {
	Debugf("hello world: %s", "cotton")
}

func TestInfo(t *testing.T) {
	Info("hello world", zap.String("name", "cotton"))
}

func TestInfof(t *testing.T) {
	Infof("hello world: %s", "cotton")
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
