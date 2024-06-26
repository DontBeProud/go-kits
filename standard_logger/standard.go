package standard_logger

import (
	"fmt"
	"github.com/DontBeProud/go-kits/error_ex"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// /RootDir/DirName/所属时间周期/ServiceName_细分时间周期.log

// StandardLoggerConfig 标准化日志配置信息
type StandardLoggerConfig struct {
	RootDir         string                 // 根目录
	Level           zapcore.Level          // 日志级别
	StackTraceLevel zapcore.LevelEnabler   // 栈追踪级别
	DirName         *string                // 文件夹名称, 若为空则默认为进程名
	RotationTime    *time.Duration         // 日志文件分割周期, 若为空则默认按小时分割
	MaxAge          *time.Duration         // 日志文件存活周期, 若为空则默认90天
	EncoderCfg      *zapcore.EncoderConfig // 编码配置, 若为空则使用默认配置
	Options         []zap.Option           // 额外的自定义选项
}

// NewStandardLoggerConfig 创建标准化日志配置
func NewStandardLoggerConfig(rootDir string, dirName *string, level zapcore.Level, rotationTime *time.Duration,
	maxAge *time.Duration, stackTraceLevel zapcore.LevelEnabler, encoderCfg *zapcore.EncoderConfig, options []zap.Option) *StandardLoggerConfig {
	return &StandardLoggerConfig{
		RootDir:         rootDir,
		Level:           level,
		StackTraceLevel: stackTraceLevel,
		DirName:         dirName,
		RotationTime:    rotationTime,
		MaxAge:          maxAge,
		EncoderCfg:      encoderCfg,
		Options:         options,
	}
}

// NewStandardLogger 创建标准的logger
// extraCallerSkip: 额外的调用栈层级过滤值
func NewStandardLogger(cfg *StandardLoggerConfig, serviceName string, extraCallerSkip *uint) (*zap.Logger, error) {
	if cfg == nil {
		return nil, error_ex.NewErrorExWithFuncNamePrefix(0, "logger config == nil")
	}

	return cfg.NewLogger(serviceName, extraCallerSkip)
}

func (cfg *StandardLoggerConfig) NewLogger(serviceName string, extraCallerSkip *uint) (*zap.Logger, error) {
	logDir := ""
	procName := strings.TrimRight(filepath.Base(os.Args[0]), ".exe") // 兼容windows

	if cfg.DirName != nil {
		logDir = *cfg.DirName
	} else {
		logDir = filepath.Join(cfg.RootDir, procName)
	}

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
	if len(cfg.Options) > 0 {
		options = append(options, cfg.Options...)
	}
	if extraCallerSkip != nil {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(int(*extraCallerSkip+1)))
	}
	if cfg.StackTraceLevel != nil {
		options = append(options, zap.AddStacktrace(cfg.StackTraceLevel))
	}

	return zap.New(zapcore.NewTee([]zapcore.Core{
		zapcore.NewCore(defaultEncoder, zapcore.AddSync(fileWriter), cfg.Level),
		zapcore.NewCore(defaultEncoder, zapcore.AddSync(os.Stdout), cfg.Level),
	}...), options...), nil
}

var (
	defaultEncoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stack",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("[2006-01-02 15:04:05.000 MST]"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)
