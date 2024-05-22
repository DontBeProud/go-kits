package standard_logger

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

func parseDurationPb(raw *durationpb.Duration) *time.Duration {
	if raw == nil {
		return nil
	}
	_duration := raw.AsDuration()
	return &_duration
}
