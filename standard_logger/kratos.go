package standard_logger

import (
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

// KratosLogger 兼容kratos的ZapLog
type KratosLogger struct {
	*zap.Logger
}

func NewKratosLogger(l *zap.Logger) (*KratosLogger, error) {
	if l == nil {
		return nil, errors.New("NewKratosLogger: logger is nil")
	}
	return &KratosLogger{l}, nil
}

func (l *KratosLogger) Log(level log.Level, kv ...interface{}) error {
	if len(kv) == 0 || len(kv)%2 != 0 {
		l.Logger.Warn(fmt.Sprint("kv must appear in pairs: ", kv))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(kv); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(kv[i]), kv[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.Logger.Debug("", data...)
	case log.LevelInfo:
		l.Logger.Info("", data...)
	case log.LevelWarn:
		l.Logger.Warn("", data...)
	case log.LevelError:
		l.Logger.Error("", data...)
	case log.LevelFatal:
		l.Logger.Fatal("", data...)
	}
	return nil
}
