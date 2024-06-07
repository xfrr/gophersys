package bus

import "context"

type HandlerMiddleware func(Handler[any, any]) Handler[any, any]

type middlewareHandler struct {
	middlewares []HandlerMiddleware
	handler     Handler[any, any]
}

func (mh *middlewareHandler) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	h := mh.handler
	for i := len(mh.middlewares) - 1; i >= 0; i-- {
		h = mh.middlewares[i](h)
	}
	return h.Handle(ctx, msg)
}

func applyMiddlewares(handler Handler[any, any], middlewares []HandlerMiddleware) Handler[any, any] {
	return &middlewareHandler{
		middlewares: middlewares,
		handler:     handler,
	}
}

// Chain creates a new middleware chain.
func Chain(middlewares ...HandlerMiddleware) HandlerMiddleware {
	return func(next Handler[any, any]) Handler[any, any] {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
