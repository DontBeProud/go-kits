package error_ex

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewErrorEx,
	NewErrorExWithFuncNamePrefix,
	NewErrorExWithPrefix,
	NewErrorExWithSuffix,
	NewErrorGroupWithContext,
)
