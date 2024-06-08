package gophergrpc

import (
	"context"
	"fmt"

	"github.com/xfrr/gophersys/pkg/bus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gopherspb "github.com/xfrr/gophersys/grpc/proto-gen-go/gopher/v1"
	gophercmd "github.com/xfrr/gophersys/internal/commands"
	gopherqry "github.com/xfrr/gophersys/internal/queries"
)

var _ gopherspb.GophersManagerServer = (*Service)(nil)

type Service struct {
	gopherspb.UnsafeGophersManagerServer

	// cmdbus is the command dispatcher.
	cmdbus bus.Bus

	// querybus is the query dispatcher.
	querybus bus.Bus
}

func NewService(cmdbus bus.Bus, querybus bus.Bus) Service {
	return Service{
		cmdbus:   cmdbus,
		querybus: querybus,
	}
}

// CreateGopher implements gopherspb.GophersManagerServer.
func (s Service) Create(ctx context.Context, req *gopherspb.CreateGopherRequest) (*gopherspb.CreateGopherResponse, error) {
	_, err := s.cmdbus.Dispatch(context.Background(), gophercmd.CreateGopherCommand{
		ID:       req.GetId(),
		Name:     req.GetName(),
		Username: req.GetUsername(),
		Status:   s.unmarshallStatus(req.GetStatus()),
		Metadata: req.GetMetadata().AsMap(),
	})
	if err != nil {
		fmt.Println(err)
		return nil, toGrpcError(err)
	}

	return &gopherspb.CreateGopherResponse{
		Id: req.GetId(),
	}, nil
}

func (s Service) Delete(ctx context.Context, req *gopherspb.DeleteGopherRequest) (*gopherspb.DeleteGopherResponse, error) {
	_, err := s.cmdbus.Dispatch(context.Background(), gophercmd.DeleteGopherCommand{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, toGrpcError(err)
	}
	return nil, nil
}

func (s Service) Get(ctx context.Context, req *gopherspb.GetGopherRequest) (*gopherspb.GetGopherResponse, error) {
	response, err := s.querybus.Dispatch(context.Background(), gopherqry.GetGopherQuery{
		GopherID: req.GetId(),
	})
	if err != nil {
		return nil, toGrpcError(err)
	}

	result, ok := response.(gopherqry.GetGopherQueryResult)
	if !ok {
		return nil, status.Errorf(codes.Internal, "unexpected query result type %T", response)
	}

	mstruct, err := structpb.NewStruct(result.Metadata)
	if err != nil {
		return nil, toGrpcError(err)
	}

	return &gopherspb.GetGopherResponse{
		Gopher: &gopherspb.Gopher{
			Id:        result.ID,
			Name:      result.Name,
			Username:  result.Username,
			Status:    s.parseStatus(result.Status),
			Metadata:  mstruct,
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}

func (s Service) List(context.Context, *gopherspb.ListGophersRequest) (*gopherspb.ListGophersResponse, error) {
	response, err := s.querybus.Dispatch(context.Background(), gopherqry.GetGophersQuery{})
	if err != nil {
		return nil, toGrpcError(err)
	}

	result, ok := response.(gopherqry.GetGophersQueryResult)
	if !ok {
		return nil, status.Errorf(codes.Internal, "unexpected query result type %T", response)
	}

	var gophers []*gopherspb.Gopher
	for _, g := range result.Gophers {
		mstruct, err := structpb.NewStruct(g.Metadata)
		if err != nil {
			return nil, toGrpcError(err)
		}

		gophers = append(gophers, &gopherspb.Gopher{
			Id:        g.ID,
			Name:      g.Name,
			Username:  g.Username,
			Status:    s.parseStatus(g.Status),
			Metadata:  mstruct,
			CreatedAt: timestamppb.New(g.CreatedAt),
			UpdatedAt: timestamppb.New(g.UpdatedAt),
		})
	}

	return &gopherspb.ListGophersResponse{
		Gophers: gophers,
	}, nil
}

func (s Service) Update(ctx context.Context, req *gopherspb.UpdateGopherRequest) (*gopherspb.UpdateGopherResponse, error) {
	_, err := s.cmdbus.Dispatch(ctx, gophercmd.UpdateGopherCommand{
		ID:       req.GetId(),
		Name:     req.GetName(),
		Username: req.GetUsername(),
		Status:   s.unmarshallStatus(req.GetStatus()),
		Metadata: req.GetMetadata().AsMap(),
	})
	if err != nil {
		return nil, toGrpcError(err)
	}

	return &gopherspb.UpdateGopherResponse{}, nil
}

func (s Service) parseStatus(status string) gopherspb.Status {
	switch status {
	case "active":
		return gopherspb.Status_ACTIVE
	case "inactive":
		return gopherspb.Status_INACTIVE
	case "suspended":
		return gopherspb.Status_SUSPENDED
	case "deleted":
		return gopherspb.Status_DELETED
	default:
		return gopherspb.Status_STATUS_UNSPECIFIED
	}
}

func (s Service) unmarshallStatus(statuspb gopherspb.Status) string {
	switch statuspb {
	case gopherspb.Status_ACTIVE:
		return "active"
	case gopherspb.Status_INACTIVE:
		return "inactive"
	case gopherspb.Status_SUSPENDED:
		return "suspended"
	case gopherspb.Status_DELETED:
		return "deleted"
	default:
		return "unknown"
	}
}
