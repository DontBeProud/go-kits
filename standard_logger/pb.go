package standard_logger

import (
	"github.com/DontBeProud/go-kits/error_ex"
	pb "github.com/DontBeProud/go-kits/standard_logger/standard_logger_pb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewStandardLoggerConfigWithPb 基于PB结构体创建标准化日志配置
func NewStandardLoggerConfigWithPb(baseCfg *pb.LoggerConfig, encoderCfg *zapcore.EncoderConfig, options []zap.Option) (*StandardLoggerConfig, error) {
	if baseCfg == nil {
		return nil, error_ex.NewErrorExWithFuncNamePrefix(0, "baseCfg == nil")
	}

	rootDir := baseCfg.RootDir
	dirName := baseCfg.DirName
	level := zapcore.Level(baseCfg.LogLevel.Number() - 1)
	rotationTime := parseDurationPb(baseCfg.RotationTime)
	maxAge := parseDurationPb(baseCfg.MaxAge)
	var stackTraceLevel zapcore.LevelEnabler
	if baseCfg.StackTraceLevel != nil {
		_level := zapcore.Level(zapcore.Level(baseCfg.StackTraceLevel.Number() - 1))
		stackTraceLevel = &_level
	}

	return NewStandardLoggerConfig(rootDir, dirName, level, rotationTime, maxAge, stackTraceLevel, encoderCfg, options), nil
}

// NewStandardLoggerWithPb 基于PB结构体创建标准化日志
func NewStandardLoggerWithPb(baseCfg *pb.LoggerConfig, serviceName string, extraCallerSkip *uint,
	encoderCfg *zapcore.EncoderConfig, options []zap.Option) (*zap.Logger, error) {
	cfg, err := NewStandardLoggerConfigWithPb(baseCfg, encoderCfg, options)
	if err != nil {
		return nil, err
	}

	return cfg.NewLogger(serviceName, extraCallerSkip)
}
