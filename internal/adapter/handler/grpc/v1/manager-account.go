package grpc_h

import (
	"context"
	pb "go-hex-arch/internal/adapter/handler/grpc/proto/manager-account"
	"go-hex-arch/internal/core/port"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ManagerAccountHandler struct {
	svc port.ManagerAccountService
	pb.UnimplementedManagerAccountServer
}

// NewCategoryHandler creates a new ManagerAccountHandler instance
func NewManagerAccountHandler(svc port.ManagerAccountService) *ManagerAccountHandler {
	return &ManagerAccountHandler{
		svc: svc,
	}
}

// SayHello implements helloworld.GreeterServer
func (s *ManagerAccountHandler) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *ManagerAccountHandler) Get(ctx context.Context, in *pb.GetManagerRequest) (*pb.GetManagerResponse, error) {
	id, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid type field: Id")

	}
	res, err := s.svc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetManagerResponse{
		Id:    res.ID.String(),
		Name:  res.Name,
		Phone: res.Phone,
	}, nil
}
