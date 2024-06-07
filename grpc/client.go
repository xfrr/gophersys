package gophergrpc

import (
	"google.golang.org/grpc"

	gopherspb "github.com/xfrr/gophersys/grpc/proto-gen-go/gopher/v1"
)

// ClientAddress represents the address of a gRPC Client.
type ClientAddress string

// String returns the string representation of the ClientAddress.
func (c ClientAddress) String() string {
	return string(c)
}

// Client represents a gRPC Client for GophersManager gRPC service.
type Client struct {
	gopherspb.GophersManagerClient
}

// NewClient creates a new gRPC Client for GophersManager gRPC service
// and returns it.
// It returns an error if the client could not be created.
func NewClient(addr ClientAddress, opts ...grpc.DialOption) (*Client, error) {
	cc, err := newClientConn(addr.String(), opts...)
	if err != nil {
		return nil, err
	}

	return &Client{gopherspb.NewGophersManagerClient(cc)}, nil
}

func newClientConn(addr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.NewClient(
		addr,
		opts...,
	)
}
