package commands

import (
	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/bus"
	"github.com/xfrr/gophersys/pkg/command"
	"github.com/xfrr/gophersys/pkg/logger"
)

func NewBus(repo gopher.Repository, log logger.Logger) *bus.InMemoryMessageBus {
	cmdbus := bus.NewInMemoryMessageBus()
	cmdbus.RegisterHandler(
		CreateGopherCommand{},
		command.NewHandler(NewCreateGopherCommandHandler(repo).Handle),
	)
	cmdbus.RegisterHandler(
		DeleteGopherCommand{},
		command.NewHandler(NewDeleteGopherCommandHandler(repo).Handle),
	)
	cmdbus.RegisterHandler(
		UpdateGopherCommand{},
		command.NewHandler(NewUpdateGopherCommandHandler(repo).Handle),
	)
	cmdbus.RegisterHandler(
		DeleteGopherCommand{},
		command.NewHandler(NewDeleteGopherCommandHandler(repo).Handle),
	)

	cmdbus.Use(bus.LoggingMiddleware(log))
	cmdbus.Use(bus.RecoverMiddleware(log))
	return cmdbus
}
