package commands

import (
	"context"

	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/command"
)

var _ command.Handler[DeleteGopherCommand, any] = (*DeleteGopherCommandHandler)(nil)

type DeleteGopherCommand struct {
	ID string
}

type DeleteGopherCommandHandler struct {
	gophers gopher.Repository
}

func (h DeleteGopherCommandHandler) Handle(ctx context.Context, command DeleteGopherCommand) (interface{}, error) {
	return nil, h.gophers.Delete(ctx, gopher.ID(command.ID))
}

// NewDeleteGopherCommandHandler creates a new DeleteGopherCommandHandler.
func NewDeleteGopherCommandHandler(gophers gopher.Repository) DeleteGopherCommandHandler {
	return DeleteGopherCommandHandler{
		gophers: gophers,
	}
}
