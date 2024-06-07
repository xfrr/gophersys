package bus

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

type ErrHandlerNotFound struct {
	MessageName string
}

func (e ErrHandlerNotFound) Error() string {
	return "handler not found for message: " + e.MessageName
}

type InMemoryMessageBus struct {
	handlers sync.Map
}

func NewInMemoryMessageBus() *InMemoryMessageBus {
	return &InMemoryMessageBus{
		handlers: sync.Map{},
	}
}

func (bus *InMemoryMessageBus) RegisterHandler(msg any, handler Handler[any, any]) {
	bus.handlers.Store(typeOf(msg), handler)
}

func (bus *InMemoryMessageBus) Dispatch(ctx context.Context, msg any) (interface{}, error) {
	handler, exists := bus.handlers.Load(typeOf(msg))
	if !exists {
		return nil, errors.New("no handler registered for message: " + typeOf(msg))
	}

	return handler.(Handler[any, any]).Handle(ctx, msg)
}

// Use registers a message handler middleware to the message bus.
// The middleware is applied to all messages dispatched through the bus.
// The middleware is applied in the order it was registered.
func (bus *InMemoryMessageBus) Use(middleware HandlerMiddleware) {
	bus.handlers.Range(func(key, value any) bool {
		handler, ok := value.(Handler[any, any])
		if !ok {
			return true
		}

		handlerWithMiddleware := applyMiddlewares(handler, []HandlerMiddleware{middleware})
		bus.handlers.Store(key, handlerWithMiddleware)
		return true
	})
}

func typeOf(msg any) string {
	return reflect.TypeOf(msg).String()
}
