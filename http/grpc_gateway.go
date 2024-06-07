package gopherhttp

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

// GrpcGateway represents a HTTP Server that serves the gRPC Gateway.
type GrpcGateway struct {
	// Server is the underlying HTTP server.
	Server *http.Server

	// Gateway is the gRPC Gateway.
	Gateway *runtime.ServeMux

	// Options are the options for the gRPC Gateway.
	Options []runtime.ServeMuxOption

	// Register is the function to register the gRPC Gateway.
	Register func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

var (
	defaultOptions = []runtime.ServeMuxOption{
		runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(runtime.DefaultHeaderMatcher),

		// json marshaller
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard, &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
	}
)

// NewGrpcGateway creates a new HTTP Server that serves the gRPC Gateway.
func NewGrpcGateway(addr string, opts ...runtime.ServeMuxOption) *GrpcGateway {
	return &GrpcGateway{
		Server:  &http.Server{Addr: addr},
		Gateway: runtime.NewServeMux(opts...),
		Options: append(defaultOptions, opts...),
	}
}

// Start starts the HTTP Server.
func (g *GrpcGateway) Start() error {
	g.Server.Handler = g.Gateway
	return g.Server.ListenAndServe()
}

// Stop stops the HTTP Server.
func (g *GrpcGateway) Stop() error {
	return g.Server.Shutdown(context.Background())
}
