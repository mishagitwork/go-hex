package repository

import (
	"go-hex-arch/internal/adapter/storage/mongo/repository/manager_account"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewManagerAccountRepository(db *mongo.Database) *manager_account.ManagerAccountRepository {
	return manager_account.New(db)
}
