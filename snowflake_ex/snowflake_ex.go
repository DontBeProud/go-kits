package snowflake_ex

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewNode,
)
