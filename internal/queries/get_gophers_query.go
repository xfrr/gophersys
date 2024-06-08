package queries

import (
	"context"
	"time"

	"github.com/xfrr/gophersys/internal/gopher"
)

// GetGophersQuery represents the query to get all gophers
type GetGophersQuery struct{}

// GetGophersQueryResult represents the result of GetGophersQuery
type GetGophersQueryResult struct {
	Gophers []GopherView
}

// GopherView represents the gopher entity
type GopherView struct {
	ID        string
	Name      string
	Username  string
	Status    string
	Metadata  map[string]any
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetGophersQueryHandler struct {
	repo gopher.Repository
}

func NewGetGophersQueryHandler(repo gopher.Repository) GetGophersQueryHandler {
	return GetGophersQueryHandler{repo: repo}
}

func (h GetGophersQueryHandler) Handle(ctx context.Context, query GetGophersQuery) (GetGophersQueryResult, error) {
	gophers, err := h.repo.Search(ctx, gopher.SearchQuery{})
	if err != nil {
		return GetGophersQueryResult{}, err
	}

	return GetGophersQueryResult{Gophers: toGophers(gophers)}, nil
}

func toGophers(gophers []*gopher.Aggregate) []GopherView {
	var res []GopherView
	for _, g := range gophers {
		res = append(res, toGopher(g))
	}

	return res
}

func toGopher(g *gopher.Aggregate) GopherView {
	return GopherView{
		ID:        g.ID().String(),
		Name:      g.Name().String(),
		Username:  g.Username().String(),
		Metadata:  g.Metadata().AsMap(),
		Status:    g.Status().String(),
		CreatedAt: g.CreatedAt(),
		UpdatedAt: g.UpdatedAt(),
	}
}
