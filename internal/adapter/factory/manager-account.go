package factory

import (
	grpc_h "go-hex-arch/internal/adapter/handler/grpc/v1"
	"go-hex-arch/internal/adapter/handler/http/v1"
	"go-hex-arch/internal/adapter/storage/database/repository"
	"go-hex-arch/internal/core/port"
	"go-hex-arch/internal/core/service"

	"gorm.io/gorm"
)

type ManagerAccountFactory struct {
	Srv     port.ManagerAccountService
	Hdl     *http.ManagerAccountHandler
	HdlGRPC *grpc_h.ManagerAccountHandler
}

// NewCategoryHandler creates a new ManagerAccountFactory instance
func NewManagerAccountFactory(db *gorm.DB) *ManagerAccountFactory {

	managerRepo := repository.NewManagerAccountRepository(db)
	managerService := service.NewCategoryService(managerRepo)
	managerHandler := http.NewManagerAccountHandler(managerService)
	managerHandlerGRPC := grpc_h.NewManagerAccountHandler(managerService)
	return &ManagerAccountFactory{
		managerService, managerHandler, managerHandlerGRPC,
	}
}
