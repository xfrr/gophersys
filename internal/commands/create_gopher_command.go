package commands

import (
	"context"

	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/command"
)

var _ command.Handler[CreateGopherCommand, any] = (*CreateGopherCommandHandler)(nil)

type CreateGopherCommand struct {
	ID       string
	Name     string
	Username string
	Status   string
	Metadata map[string]any
}

type CreateGopherCommandHandler struct {
	gophers gopher.Repository
}

func (h CreateGopherCommandHandler) Handle(ctx context.Context, cmd CreateGopherCommand) (interface{}, error) {
	ok, err := h.gophers.Exists(
		ctx,
		gopher.ByID(gopher.ID(cmd.ID)),
		gopher.ByUsername(gopher.Username(cmd.Username)),
	)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, gopher.ErrAlreadyExistsByUsernameOrID
	}

	newGopher, err := gopher.New(
		cmd.ID,
		gopher.WithName(cmd.Name),
		gopher.WithStatus(cmd.Status),
		gopher.WithUsername(cmd.Username),
		gopher.WithMetadata(gopher.ParseMetadata(cmd.Metadata)),
	)
	if err != nil {
		return nil, err
	}

	return nil, h.gophers.Save(ctx, newGopher)
}

func NewCreateGopherCommandHandler(gophers gopher.Repository) CreateGopherCommandHandler {
	return CreateGopherCommandHandler{
		gophers: gophers,
	}
}
