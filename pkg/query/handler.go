package query

import (
	"context"
	"fmt"
	"reflect"
)

type ErrInvalidMessage struct {
	msg interface{}
}

func NewErrInvalidMessage(msg interface{}) ErrInvalidMessage {
	return ErrInvalidMessage{msg: msg}
}

func (e ErrInvalidMessage) Error() string {
	return fmt.Sprintf("invalid message: %T", e.msg)
}

var _ Handler[any, any] = (*handler)(nil)

// ErrInvalidHandlerSignature is the error returned when the handler signature is invalid.
type ErrInvalidHandlerSignature struct {
	msg interface{}
	h   interface{}
}

func (e ErrInvalidHandlerSignature) Error() string {
	return fmt.Sprintf("invalid handler signature: %T(%T)", e.h, e.msg)
}

// Handler is a struct that holds the reference to the Handle method.
type Handler[C any, R any] interface {
	Handle(ctx context.Context, msg C) (R, error)
}

// HandlerFunc is a shortcut for a function that handles a msg.
type HandlerFunc func(ctx context.Context, msg any) (any, error)

// Handle calls the underlying HandlerFunc.
func (f HandlerFunc) Handle(ctx context.Context, msg any) (any, error) {
	return HandlerFunc(f)(ctx, msg)
}

// Handler is a struct that holds the reference to the Handle method.
type handler struct {
	handleFn func(ctx context.Context, msg any) (any, error)
}

// Handle calls the Handle method.
func (h handler) Handle(ctx context.Context, msg any) (any, error) {
	return h.handleFn(ctx, msg)
}

func wrapHandlerFunc[C any, R any](fn func(ctx context.Context, msg C) (R, error)) func(ctx context.Context, msg any) (any, error) {
	return func(ctx context.Context, msg any) (any, error) {
		c, ok := msg.(C)
		if !ok {
			return nil, NewErrInvalidMessage(msg)
		}

		r, err := fn(ctx, c)
		return r, err
	}
}

// isValidHandlerFunc checks if the function is a valid handler.
func isValidHandlerFunc(fn any) bool {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return false
	}

	ft := reflect.TypeOf(fn)
	if ft.NumIn() != 2 {
		return false
	}

	if ft.NumOut() != 2 {
		return false
	}

	if ft.Out(1).String() != "error" {
		return false
	}

	return true
}

// NewHandler creates a new Handler.
func NewHandler[C any, R any](handleFn func(ctx context.Context, msg C) (R, error)) Handler[any, any] {
	if !isValidHandlerFunc(handleFn) {
		panic("handleFn must be a function")
	}

	return handler{
		handleFn: wrapHandlerFunc(handleFn),
	}
}
