package gophermongo

import (
	"time"

	"github.com/xfrr/gophersys/internal/gopher"
)

type gopherDTO struct {
	ID        string         `bson:"_id"`
	Name      string         `bson:"name"`
	Username  string         `bson:"username"`
	Status    string         `bson:"status"`
	Metadata  map[string]any `bson:"metadata"`
	CreatedAt time.Time      `bson:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at"`
}

func gopherToDTO(g gopher.Aggregate) gopherDTO {
	return gopherDTO{
		ID:        g.ID().String(),
		Name:      g.Name().String(),
		Username:  g.Username().String(),
		Status:    g.Status().String(),
		Metadata:  g.Metadata().AsMap(),
		CreatedAt: g.CreatedAt(),
		UpdatedAt: g.UpdatedAt(),
	}
}

func dtoToGopher(dto gopherDTO) (*gopher.Aggregate, error) {
	return gopher.New(
		dto.ID,
		gopher.WithName(dto.Name),
		gopher.WithUsername(dto.Username),
		gopher.WithStatus(dto.Status),
		gopher.WithMetadata(gopher.ParseMetadata(dto.Metadata)),
		gopher.WithCreatedAt(dto.CreatedAt),
		gopher.WithUpdatedAt(dto.UpdatedAt),
	)
}
