// gopher package represents the domain layer of the Gophers application.
package gopher

import (
	"time"
)

type gopher struct {
	id        ID
	name      Name
	username  Username
	status    Status
	metadata  Metadata
	createdAt time.Time
	updateAt  time.Time
}

// Aggregate represents a gopher aggregate
type Aggregate struct {
	gopher
}

// ID returns the Gopher ID
func (g Aggregate) ID() ID {
	return g.id
}

// Name returns the Gopher name
func (g Aggregate) Name() Name {
	return g.name
}

// Username returns the Gopher username
func (g Aggregate) Username() Username {
	return g.username
}

// Status returns the Gopher status
func (g Aggregate) Status() Status {
	return g.status
}

// Metadata returns the Gopher metadata
func (g Aggregate) Metadata() Metadata {
	return g.metadata
}

// CreatedAt returns the Gopher creation time
func (g Aggregate) CreatedAt() time.Time {
	return g.createdAt
}

// UpdatedAt returns the Gopher update time
func (g Aggregate) UpdatedAt() time.Time {
	return g.updateAt
}

// SetName sets the name for the Gopher entity
func (g *Aggregate) SetName(name string) {
	g.name = Name(name)
}

// SetMetadata sets the metadata for the Gopher entity
func (g *Aggregate) MergeMetadata(metadata Metadata) {
	g.metadata.merge(metadata)
}

// ChangeUsername changes the username for the Gopher entity
func (g *Aggregate) ChangeUsername(username string) error {
	if !g.username.IsValid() {
		return ErrInvalidUsername
	}

	g.username = Username(username)
	return nil
}

// ChangeStatus changes the status for the Gopher entity
func (g *Aggregate) ChangeStatus(status Status) error {
	if !g.status.IsValid() {
		return ErrInvalidStatus
	}

	g.status = status
	return nil
}

// validate checks if the Gopher entity is valid
func (g Aggregate) validate() error {
	if !g.id.IsValid() {
		return ErrInvalidID
	}

	if !g.username.IsValid() {
		return ErrInvalidUsername
	}

	if !g.status.IsValid() {
		return ErrInvalidStatus
	}

	return nil
}

// New creates a new Gopher aggregate
func New(id string, modifiers ...Modifier) (*Aggregate, error) {
	g := &Aggregate{
		gopher: gopher{
			id:        ID(id),
			metadata:  Metadata{},
			status:    StatusActive,
			createdAt: time.Now(),
			updateAt:  time.Now(),
		},
	}

	for _, m := range modifiers {
		m(g)
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return g, nil
}

// Modifier represents a variadic function for Gopher creation
type Modifier func(*Aggregate)

// WithName sets the name for the Gopher entity
func WithName(name string) Modifier {
	return func(g *Aggregate) {
		if Name(name).IsValid() {
			g.name = Name(name)
		}
	}
}

// WithUsername sets the username for the Gopher entity
func WithUsername(username string) Modifier {
	return func(g *Aggregate) {
		if Username(username).IsValid() {
			g.username = Username(username)
		}
	}
}

// WithMetadata sets the metadata for the Gopher entity
func WithMetadata(metadata Metadata) Modifier {
	return func(g *Aggregate) {
		if !metadata.IsEmpty() {
			g.metadata = metadata
		}
	}
}

// WithStatus sets the status for the Gopher entity
func WithStatus(statusStr string) Modifier {
	return func(g *Aggregate) {
		status := StatusFromString(statusStr)
		if status.IsValid() {
			g.status = status
		}
	}
}

// WithCreatedAt sets the creation time for the Gopher entity
func WithCreatedAt(t time.Time) Modifier {
	return func(g *Aggregate) {
		if g.createdAt.IsZero() {
			g.createdAt = t
		}
	}
}

// WithUpdatedAt sets the update time for the Gopher entity
func WithUpdatedAt(t time.Time) Modifier {
	return func(g *Aggregate) {
		if g.updateAt.IsZero() {
			g.updateAt = t
		}
	}
}
