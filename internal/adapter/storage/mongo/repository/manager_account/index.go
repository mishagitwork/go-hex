package manager_account

import "go.mongodb.org/mongo-driver/mongo"

type ManagerAccountRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) *ManagerAccountRepository {
	return &ManagerAccountRepository{db: db}
}
