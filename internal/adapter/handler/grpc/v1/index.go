package grpc_h

import (
	"errors"
	"fmt"

	pb "go-hex-arch/internal/adapter/handler/grpc/proto/manager-account"
	"go-hex-arch/internal/configs"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Router is a wrapper for HTTP router
type Router struct {
	*grpc.Server
}

func NewRouter(
	config *configs.App,
	managerAccountHandler *ManagerAccountHandler,

) *Router {

	router := grpc.NewServer()
	pb.RegisterManagerAccountServer(router, managerAccountHandler)

	reflection.Register(router)

	return &Router{
		router,
	}
}

// Serve starts the GRPC server
func (r *Router) Run(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.New("failed to listen: " + err.Error())
	}
	return r.Serve(lis)
}
