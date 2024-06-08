package gophergrpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	gopherspb "github.com/xfrr/gophersys/grpc/proto-gen-go/gopher/v1"
	healtcheckpb "google.golang.org/grpc/health/grpc_health_v1"
	reflectionpb "google.golang.org/grpc/reflection"
)

var (
	defaultServerOptions = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 4), // 4 MB
		grpc.MaxSendMsgSize(1024 * 1024 * 4), // 4 MB
	}
)

// ServerAddress represents the address of a gRPC Server.
type ServerAddress string

// String returns the string representation of the ServerAddress.
func (s ServerAddress) String() string {
	return string(s)
}

// MakeServerAddress creates a new ServerAddress from the given address.
func MakeServerAddress(host, port string) ServerAddress {
	return ServerAddress(net.JoinHostPort(host, port))
}

// Server represents a gRPC Server for GophersManager gRPC service.
type Server struct {
	// Address is the address of the gRPC Server.
	Address ServerAddress

	// srv is the underlying gRPC Server.
	srv *grpc.Server

	// svc implements the GophersManagerServer interface.
	svc gopherspb.GophersManagerServer
}

// Start starts the gRPC Server.
func (s Server) Start() error {
	lis, err := net.Listen("tcp", s.Address.String())
	if err != nil {
		return err
	}

	return s.srv.Serve(lis)
}

// GracefulStop stops the gRPC Server gracefully.
func (s Server) GracefulStop() {
	s.srv.GracefulStop()
}

// NewServer creates a new gRPC Server for GophersManager gRPC service
// and returns it.
// It returns an error if the server could not be created.
func NewServer(
	port string,
	svc gopherspb.GophersManagerServer,
	opts ...grpc.ServerOption,
) *Server {
	srv := &Server{
		Address: MakeServerAddress("localhost", port),
		svc:     svc,
		srv:     grpc.NewServer(append(defaultServerOptions, opts...)...),
	}
	gopherspb.RegisterGophersManagerServer(srv.srv, srv.svc)
	healtcheckpb.RegisterHealthServer(srv.srv, health.NewServer())
	reflectionpb.Register(srv.srv)
	return srv
}
