package utils

import (
	"go-hex-arch/internal/core/dto"

	"gorm.io/gorm"
)

// Paginate ...
func Paginate(q dto.Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if q.Page == 0 {
			q.Page = 1
		}

		switch {
		case q.Limit > 100:
			q.Limit = 100
		case q.Limit == 0:
			q.Limit = 10
		case q.Limit < 0:
			q.Limit = -1
		}

		offset := (q.Page - 1) * q.Limit
		return db.Offset(offset).Limit(q.Limit)
	}
}
