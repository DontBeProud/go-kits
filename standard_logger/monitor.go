package standard_logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
	"time"
)

// MonitorLogger 监控日志对象
type MonitorLogger interface {
	// ExecuteMonitorTask 执行监控任务
	// interval: 轮询间隔
	// callback: 自定义回调函数
	ExecuteMonitorTask(ctx context.Context, interval time.Duration, callback func(*runtime.MemStats, *zap.Logger))
	// GetLogger 获取底层日志对象
	GetLogger() *zap.Logger
}

// NewMonitorLogger 创建标准化监控日志对象
func NewMonitorLogger(cfg *StandardLoggerConfig, serviceName string) (MonitorLogger, error) {
	loggerObj, err := NewStandardLogger(cfg, serviceName, nil)
	if err != nil {
		return nil, err
	}
	return &monitorLogger{Logger: loggerObj}, nil
}

// monitorLogger 标准化监控日志对象
type monitorLogger struct {
	*zap.Logger
}

func (l *monitorLogger) GetLogger() *zap.Logger {
	return l.Logger
}

// ExecuteMonitorTask 执行监控任务
// interval: 轮询间隔
// callback: 自定义回调函数
func (l *monitorLogger) ExecuteMonitorTask(ctx context.Context, interval time.Duration, callback func(*runtime.MemStats, *zap.Logger)) {
	defer func() {
		if err := recover(); err != nil {
			l.Error(fmt.Sprintf("monitor err:%v;stack:%v", err, string(debug.Stack())))
		}
	}()

	memStatus := runtime.MemStats{}
	tick := time.NewTicker(interval)
	for {
		runtime.ReadMemStats(&memStatus)

		l.Info(fmt.Sprintf("current Goroutine:%-6d Heap:%.2fMb Stack:%.2fMb",
			runtime.NumGoroutine(),
			float64(memStatus.HeapInuse)/1024.0/1024.0,
			float64(memStatus.StackInuse)/1024.0/1024.0,
		))

		if callback != nil {
			callback(&memStatus, l.Logger)
		}

		select {
		case <-ctx.Done():
			break
		case <-tick.C:
			continue
		}
	}
}
