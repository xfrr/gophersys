package mongoutils

import (
	"context"
)

type Migration interface {
	Name() string
	Up(ctx context.Context) error
	Down(ctx context.Context) error
}

type Migrator interface {
	ApplyMigrations(ctx context.Context) error
	RemoveMigrations(ctx context.Context) error
}
