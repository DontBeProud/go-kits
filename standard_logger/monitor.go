package standard_logger

import (
	"context"
	"fmt"
	"github.com/DontBeProud/go-kits/error_ex"
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
	"time"
)

// GlobalMonitorLogger 全局监控日志对象
type GlobalMonitorLogger interface {
	// ExecuteGlobalMonitorTask 执行全局监控任务
	// interval: 轮询间隔
	// callback: 自定义回调函数
	ExecuteGlobalMonitorTask(ctx context.Context, interval time.Duration, callback func(*runtime.MemStats, *zap.Logger))
	// GetLogger 获取底层日志对象
	GetLogger() *zap.Logger
}

// NewMonitorLogger 创建标准化监控日志对象
func NewMonitorLogger(cfg *StandardLoggerConfig) (GlobalMonitorLogger, error_ex.ErrorEx) {
	const serviceName = "global_monitor"
	loggerObj, err := NewStandardLogger(cfg, serviceName, nil)
	if err != nil {
		return nil, err
	}
	return &globalMonitorLogger{Logger: loggerObj}, nil
}

// globalMonitorLogger 标准化监控日志对象
type globalMonitorLogger struct {
	*zap.Logger
}

func (l *globalMonitorLogger) GetLogger() *zap.Logger {
	return l.Logger
}

// ExecuteGlobalMonitorTask 执行全局监控任务
// interval: 轮询间隔
// callback: 自定义回调函数
func (l *globalMonitorLogger) ExecuteGlobalMonitorTask(ctx context.Context, interval time.Duration, callback func(*runtime.MemStats, *zap.Logger)) {
	defer func() {
		if err := recover(); err != nil {
			l.Error(fmt.Sprintf("monitor err:%v;stack:%v", err, string(debug.Stack())))
		}
	}()

	tick := time.NewTicker(interval)
	for {
		memStatus := &runtime.MemStats{}
		runtime.ReadMemStats(memStatus)

		l.Info(fmt.Sprintf("current Goroutine:%-6d Heap:%.2fMb Stack:%.2fMb",
			runtime.NumGoroutine(),
			float64(memStatus.HeapInuse)/1024.0/1024.0,
			float64(memStatus.StackInuse)/1024.0/1024.0,
		))

		if callback != nil {
			callback(memStatus, l.Logger)
		}

		select {
		case <-ctx.Done():
			break
		case <-tick.C:
		}
	}
}
