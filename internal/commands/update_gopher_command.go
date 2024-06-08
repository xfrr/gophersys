package commands

import (
	"context"

	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/command"
)

var _ command.Handler[UpdateGopherCommand, any] = (*UpdateGopherCommandHandler)(nil)

type UpdateGopherCommand struct {
	ID       string
	Name     string
	Username string
	Status   string
	Metadata map[string]any
}

type UpdateGopherCommandHandler struct {
	gophers gopher.Repository
}

func (h UpdateGopherCommandHandler) Handle(ctx context.Context, cmd UpdateGopherCommand) (interface{}, error) {
	foundGopher, err := h.gophers.Get(
		ctx,
		gopher.ByID(gopher.ID(cmd.ID)),
	)
	if err != nil {
		return nil, err
	}

	err = foundGopher.Update(
		gopher.Name(cmd.Name),
		gopher.Username(cmd.Username),
		gopher.ParseStatus(cmd.Status),
		gopher.ParseMetadata(cmd.Metadata),
	)
	if err != nil {
		return nil, err
	}

	return nil, h.gophers.Save(ctx, foundGopher)
}

func NewUpdateGopherCommandHandler(gophers gopher.Repository) UpdateGopherCommandHandler {
	return UpdateGopherCommandHandler{
		gophers: gophers,
	}
}
