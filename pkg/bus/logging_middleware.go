package bus

import (
	"context"
	"fmt"

	"github.com/xfrr/gophersys/pkg/logger"
)

// LoggingMiddleware is a simple message handler middleware
// that logs the message dispatched.
func LoggingMiddleware(logger logger.Logger) HandlerMiddleware {
	return HandlerMiddleware(func(next Handler[any, any]) Handler[any, any] {
		return HandlerFunc(func(ctx context.Context, msg interface{}) (interface{}, error) {
			logger.Info().
				Str("name", fmt.Sprintf("%T", msg)).
				Str("component", "bus").
				Msgf("message dispatched")
			return next.Handle(ctx, msg)
		})
	})
}
