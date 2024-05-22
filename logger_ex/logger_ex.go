package logger_ex

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewStandardLogger,
	NewStandardLoggerConfig,
)
