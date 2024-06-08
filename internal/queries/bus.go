package queries

import (
	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/bus"
	"github.com/xfrr/gophersys/pkg/logger"
	"github.com/xfrr/gophersys/pkg/query"
)

// NewBus creates a new query bus with the registered handlers.
func NewBus(repo gopher.Repository, log logger.Logger) *bus.InMemoryMessageBus {
	cmdbus := bus.NewInMemoryMessageBus()
	cmdbus.RegisterHandler(
		GetGophersQuery{},
		query.NewHandler(NewGetGophersQueryHandler(repo).Handle),
	)
	cmdbus.RegisterHandler(
		GetGopherQuery{},
		query.NewHandler(NewGetGopherQueryHandler(repo).Handle),
	)

	cmdbus.Use(bus.LoggingMiddleware(log))
	cmdbus.Use(bus.RecoverMiddleware(log))
	return cmdbus
}
