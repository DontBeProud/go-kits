package standard_logger

import (
	"context"
	"errors"
	"fmt"
	"github.com/DontBeProud/go-kits/error_ex"
	"github.com/DontBeProud/go-kits/gorm_ex"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

// GormLoggerConfig GormLogger config
type GormLoggerConfig struct {
	LogLevel                      logger.LogLevel // Log level
	SlowThreshold                 time.Duration   // Slow SQL threshold
	DontIgnoreRecordNotFoundError bool            // Don't Ignore ErrRecordNotFound error for logger
	DontIgnoreKeyDuplicateError   bool            // Don't Ignore ErrKeyDuplicate error for logger
	FilterParams                  bool            // hide params when print sql
}

// NewStandardGormLogger 新建标准的gorm日志对象
func NewStandardGormLogger(loggerCore *zap.Logger, cfg *GormLoggerConfig) (logger.Interface, error_ex.ErrorEx) {
	const callerSkip = 2
	if loggerCore == nil {
		return nil, error_ex.NewErrorEx("invalid loggerCore")
	}
	if cfg == nil {
		return nil, error_ex.NewErrorEx("invalid GormLoggerConfig")
	}
	return &GormLogger{
		GormLoggerConfig: *cfg,
		loggerCore:       loggerCore.WithOptions(zap.AddCallerSkip(callerSkip)),
	}, nil
}

type GormLogger struct {
	GormLoggerConfig
	loggerCore *zap.Logger
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.loggerCore.Info(fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

// Warn print warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.loggerCore.Warn(fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

// Error print error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.loggerCore.Error(fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	var (
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil &&
		l.LogLevel >= logger.Error &&
		(!errors.Is(err, logger.ErrRecordNotFound) || l.DontIgnoreRecordNotFoundError) &&
		(!gorm_ex.IsErrorDuplicateKey(err) || l.DontIgnoreKeyDuplicateError):

		if rows == -1 {
			l.loggerCore.Error(fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.loggerCore.Error(fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.loggerCore.Warn(fmt.Sprintf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.loggerCore.Warn(fmt.Sprintf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case l.LogLevel == logger.Info:
		if rows == -1 {
			l.loggerCore.Info(fmt.Sprintf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.loggerCore.Info(fmt.Sprintf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	}
}

// ParamsFilter filter params
func (l *GormLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if l.GormLoggerConfig.FilterParams {
		return sql, nil
	}
	return sql, params
}
