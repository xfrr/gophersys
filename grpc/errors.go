package gophergrpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gophercmd "github.com/xfrr/gophersys/internal/commands"
)

func toGrpcError(err error) error {
	switch err {
	case gophercmd.ErrGopherAlreadyExists:
		return status.Error(codes.AlreadyExists, "gopher already exists")
	case gophercmd.ErrGopherNotFound:
		return status.Error(codes.NotFound, "gopher not found")
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
