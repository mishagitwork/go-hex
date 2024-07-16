package mongo

import (
	"context"
	"fmt"
	"go-hex-arch/internal/configs"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New ...
func New(config *configs.Mongo) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.Username, config.Password, config.Host, config.Port)
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = connect.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := connect.Database(config.Database)

	return db, nil
}
