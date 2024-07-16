package utils

import (
	"go-hex-arch/internal/core/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetPagination ...
func SetSort(opt *options.FindOptions, q dto.Query) {
	direction := 1

	if q.OrderBy == "" {
		return
	}

	switch q.OrderDir {
	case "asc":
		direction = 1
	case "desc":
		direction = -1
	}

	sort := bson.D{{Key: q.OrderBy, Value: direction}}
	opt.SetSort(sort)
}
