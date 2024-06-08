package gopher

import "context"

type SDK interface {
	// Dispatch dispatches the given command to the underlying command bus
	// The behavior of the command bus is implementation specific.
	// It returns an error if the command could not be dispatched.
	Dispatch(ctx context.Context, cmd interface{}) error

	// Query dispatches the given query to the underlying query bus
	// The behavior of the query bus is implementation specific.
	// It returns an error if the query could not be dispatched.
	Query(ctx context.Context, query interface{}) (interface{}, error)
}
