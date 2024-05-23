package standard_logger

import (
	"github.com/DontBeProud/go-kits/error_ex"
	pb "github.com/DontBeProud/go-kits/standard_logger/standard_logger_pb"
	"go.uber.org/zap/zapcore"
)

// NewStandardLoggerWithPb 基于PB结构体创建标准化日志配置
func NewStandardLoggerWithPb(baseCfg *pb.LoggerConfig, encoderCfg *zapcore.EncoderConfig) (*StandardLoggerConfig, error) {
	if baseCfg == nil {
		return nil, error_ex.NewErrorExWithFuncNamePrefix(0, "baseCfg == nil")
	}

	rootDir := baseCfg.RootDir
	level := zapcore.Level(baseCfg.LogLevel.Number() - 1)
	rotationTime := parseDurationPb(baseCfg.RotationTime)
	maxAge := parseDurationPb(baseCfg.MaxAge)
	var stackTraceLevel zapcore.LevelEnabler
	if baseCfg.StackTraceLevel != nil {
		_level := zapcore.Level(zapcore.Level(baseCfg.StackTraceLevel.Number() - 1))
		stackTraceLevel = &_level
	}

	return NewStandardLoggerConfig(rootDir, level, rotationTime, maxAge, stackTraceLevel, encoderCfg), nil
}
