package bus

import "context"

type Bus interface {
	Dispatcher
}

type Dispatcher interface {
	Dispatch(ctx context.Context, msg interface{}) (res interface{}, err error)
}

// Handler is a struct that holds the reference to the Handle method.
type Handler[M any, R any] interface {
	Handle(ctx context.Context, msg M) (R, error)
}

// HandlerFunc is a shortcut for a function that handles a msg.
type HandlerFunc func(ctx context.Context, msg any) (any, error)

// Handle calls the underlying HandlerFunc.
func (f HandlerFunc) Handle(ctx context.Context, msg any) (any, error) {
	return HandlerFunc(f)(ctx, msg)
}
