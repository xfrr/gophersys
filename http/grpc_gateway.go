package gopherhttp

import (
	"context"
	"net/http"
	"time"

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
		Server:  newServer(addr),
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

func newServer(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
		// ReadHeaderTimeout is the amount of time allowed to read
		// request headers. The connection's read deadline is reset
		// after reading the headers and the Handler can decide what
		// is considered too slow for the body. If ReadHeaderTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		ReadHeaderTimeout: 15 * time.Second,

		// ReadTimeout is the maximum duration for reading the entire
		// request, including the body. A zero or negative value means
		// there will be no timeout.
		//
		// Because ReadTimeout does not let Handlers make per-request
		// decisions on each request body's acceptable deadline or
		// upload rate, most users will prefer to use
		// ReadHeaderTimeout. It is valid to use them both.
		ReadTimeout: 15 * time.Second,

		// WriteTimeout is the maximum duration before timing out
		// writes of the response. It is reset whenever a new
		// request's header is read. Like ReadTimeout, it does not
		// let Handlers make decisions on a per-request basis.
		// A zero or negative value means there will be no timeout.
		WriteTimeout: 10 * time.Second,

		// IdleTimeout is the maximum amount of time to wait for the
		// next request when keep-alives are enabled. If IdleTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		IdleTimeout: 30 * time.Second,
	}
}
