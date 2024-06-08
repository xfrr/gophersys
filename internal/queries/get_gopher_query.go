package queries

import (
	"context"

	"github.com/xfrr/gophersys/internal/gopher"
)

type GetGopherQuery struct {
	GopherID string
}

type GetGopherQueryResult struct {
	GopherView
}

type GetGopherQueryHandler struct {
	repo gopher.Repository
}

func NewGetGopherQueryHandler(repo gopher.Repository) GetGopherQueryHandler {
	return GetGopherQueryHandler{repo: repo}
}

func (h GetGopherQueryHandler) Handle(ctx context.Context, query GetGopherQuery) (GetGopherQueryResult, error) {
	gopher, err := h.repo.Get(ctx, gopher.ByID(gopher.ID(query.GopherID)))
	if err != nil {
		return GetGopherQueryResult{}, err
	}

	return GetGopherQueryResult{GopherView: toGopher(gopher)}, nil
}
