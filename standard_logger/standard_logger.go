package standard_logger

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewStandardLogger,
	NewStandardLoggerConfig,
	NewStandardLoggerConfigWithPb,
	NewKratosLogger,
	NewStandardLoggerWithPb,
	NewGormLogger,
)
