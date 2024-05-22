package logger_ex

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go-kits/error_ex"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// StandardLoggerConfig 标准化日志配置信息
type StandardLoggerConfig struct {
	Level        zapcore.Level          // 日志级别
	RootDir      string                 // 根目录
	RotationTime *time.Duration         // 日志文件分割周期, 若为空则默认按小时分割
	MaxAge       *time.Duration         // 日志文件存活周期, 若为空则默认90天
	EncoderCfg   *zapcore.EncoderConfig // 编码配置, 若为空则使用默认配置
}

// NewStandardLoggerConfig 创建标准化日志配置
func NewStandardLoggerConfig(level zapcore.Level, rootDir string, rotationTime *time.Duration,
	maxAge *time.Duration, encoderCfg *zapcore.EncoderConfig) *StandardLoggerConfig {
	return &StandardLoggerConfig{
		Level:        level,
		RootDir:      rootDir,
		RotationTime: rotationTime,
		MaxAge:       maxAge,
		EncoderCfg:   encoderCfg,
	}
}

// NewStandardLogger 创建标准的logger
func NewStandardLogger(cfg *StandardLoggerConfig, serviceName string, callerSkip *uint) (*zap.Logger, error) {
	if cfg == nil {
		return nil, error_ex.NewErrorExWithFuncNamePrefix(0, "logger config == nil")
	}

	return cfg.NewLogger(serviceName, callerSkip)
}

func (cfg *StandardLoggerConfig) NewLogger(serviceName string, callerSkip *uint) (*zap.Logger, error) {
	procName := strings.TrimRight(filepath.Base(os.Args[0]), ".exe") // 兼容windows
	logDir := filepath.Join(cfg.RootDir, procName)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}
	if serviceName == "" {
		serviceName = procName
	}
	fileName := filepath.Join(logDir, "%v", serviceName+"_%v.log")

	fileTail := ""

	rotationTime := time.Hour
	if cfg.RotationTime != nil {
		rotationTime = *cfg.RotationTime
	}
	if rotationTime > time.Hour {
		fileTail = "%Y_%m_%d"
	} else if rotationTime > time.Minute*30 {
		fileTail = "%Y_%m_%d_%H"
	} else {
		fileTail = "%Y_%m_%d_%H_%M"
	}

	maxAge := 90 * 24 * time.Hour
	if cfg.MaxAge != nil {
		maxAge = *cfg.MaxAge
	}

	fileWriter, _ := rotatelogs.New(
		fmt.Sprintf(fileName, "%Y%m%d", fileTail),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	options := make([]zap.Option, 0)
	if callerSkip != nil {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(int(*callerSkip)))
	}

	return zap.New(zapcore.NewTee([]zapcore.Core{
		zapcore.NewCore(defaultEncoder, zapcore.AddSync(fileWriter), cfg.Level),
		zapcore.NewCore(defaultEncoder, zapcore.AddSync(os.Stdout), cfg.Level),
	}...), options...), nil
}

var (
	defaultEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("[2006-01-02 15:04:05.000 MST]"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)
