package utils

import (
	"go-hex-arch/internal/core/dto"

	"gorm.io/gorm"
)

// OrderBy ...
func OrderBy(q dto.Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		var direction string = "ASC"

		if q.OrderBy == "" {
			return db
		}

		switch q.OrderDir {
		case "asc":
			direction = "ASC"
		case "desc":
			direction = "DESC"
		}

		return db.Order(q.OrderBy + " " + direction)
	}
}
