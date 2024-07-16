package factory

import (
	"go-hex-arch/internal/adapter/handler/http/v1"
	"go-hex-arch/internal/adapter/storage/mongo/repository"
	"go-hex-arch/internal/core/port"
	"go-hex-arch/internal/core/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type ManagerAccountMongoFactory struct {
	Srv port.ManagerAccountService
	Hdl *http.ManagerAccountHandler
}

// NewCategoryHandler creates a new ManagerAccountFactory instance
func NewManagerAccountMongoFactory(db *mongo.Database, managerAccountService port.ManagerAccountService) *ManagerAccountMongoFactory {

	managerRepo := repository.NewManagerAccountRepository(db)
	managerService := service.NewManagerAccountMongoService(managerRepo, managerAccountService)
	managerHandler := http.NewManagerAccountHandler(managerService)
	return &ManagerAccountMongoFactory{
		managerService, managerHandler,
	}
}
