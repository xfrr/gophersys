package bus

import (
	"context"
	"fmt"

	"github.com/xfrr/gophersys/pkg/logger"
)

// RecoverMiddleware is a simple message handler middleware
// that recovers from panics.
func RecoverMiddleware(logger logger.Logger) HandlerMiddleware {
	return func(next Handler[any, any]) Handler[any, any] {
		return HandlerFunc(func(ctx context.Context, msg interface{}) (interface{}, error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Error().
						Str("name", fmt.Sprintf("%T", msg)).
						Str("component", "bus").
						Msgf("panic recovered: %v", r)
				}
			}()
			return next.Handle(ctx, msg)
		})
	}
}
