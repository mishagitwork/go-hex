package utils

import (
	"go-hex-arch/internal/core/dto"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetPagination ...
func SetPagination(opt *options.FindOptions, q dto.Query) {

	if q.Page == 0 {
		q.Page = 1
	}

	switch {
	case q.Limit > 100:
		q.Limit = 100
	case q.Limit <= 0:
		q.Limit = 10
	}

	offset := (q.Page - 1) * q.Limit
	opt.SetLimit(int64(q.Limit)).SetSkip(int64(offset))
}
