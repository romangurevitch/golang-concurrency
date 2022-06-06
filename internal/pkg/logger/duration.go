package logger

import (
	"context"
	"time"
)

const durationMillisKey = "duration_millis"

func TraceDuration(ctx context.Context, start time.Time, msg string) {
	duration := time.Since(start).Milliseconds()
	WithContext(ctx).WithField(durationMillisKey, duration).Trace(msg)
}
