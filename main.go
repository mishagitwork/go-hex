package main

import (
	"fmt"

	"go-hex-arch/internal/adapter/factory"
	"go-hex-arch/pkg/logger"

	grpc_h "go-hex-arch/internal/adapter/handler/grpc/v1"
	"go-hex-arch/internal/adapter/handler/http/v1"
	"go-hex-arch/internal/adapter/storage/database"
	"go-hex-arch/internal/adapter/storage/mongo"
	"go-hex-arch/internal/configs"
	"log"
)

var (
	config *configs.Config
)

func init() {
	var err error

	config, err = configs.New()
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	logger := logger.New()
	db, err := database.New(&config.Database)
	if err != nil {
		logger.Fatal(err)
	}
	mongo, err := mongo.New(&config.Mongo)
	if err != nil {
		logger.Fatal(err)
	}

	managerFactory := factory.NewManagerAccountFactory(db)
	managerMongoFactory := factory.NewManagerAccountMongoFactory(mongo, managerFactory.Srv)

	router := http.NewRouter(
		&config.App,
		logger,
		managerMongoFactory.Hdl,
	)

	grpcRouter := grpc_h.NewRouter(
		&config.App,
		managerFactory.HdlGRPC,
	)

	go func() {
		err = grpcRouter.Run(int(config.App.GRPCPort))
		if err != nil {
			logger.Fatal(err)
		}
		fmt.Printf("GRPC server has been successfully started on port: %d\n", config.App.GRPCPort)
	}()

	listenAddr := fmt.Sprintf(":%d", config.App.HTTPPort)
	err = router.Serve(listenAddr)
	if err != nil {
		logger.Fatal(err)
	}
}
