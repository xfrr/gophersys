package commands

import "github.com/xfrr/gophersys/internal/gopher"

// These are a copy of the errors from the gopher package
// to avoid an import from grpc package directly to the domain layer.
var (
	ErrGopherAlreadyExists = gopher.ErrAlreadyExistsByUsernameOrID
	ErrGopherNotFound      = gopher.ErrNotFound
)
