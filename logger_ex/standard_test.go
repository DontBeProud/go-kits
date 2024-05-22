package logger_ex

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewStandardLoggerConfig(t *testing.T) {
	cfg := NewStandardLoggerConfig(zapcore.InfoLevel, "", nil, nil, nil, nil)

	obj, _ := cfg.NewLogger("test", nil)
	obj.Info("aha")

	var skip uint = 0
	objWithCaller, _ := cfg.NewLogger("test-with-caller", &skip)
	objWithCaller.Info("wow")

}
