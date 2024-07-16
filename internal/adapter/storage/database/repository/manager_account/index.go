package manager_account

import (
	"go-hex-arch/internal/adapter/storage/database/models"

	"gorm.io/gorm"
)

type ManagerAccountRepository struct {
	db  *gorm.DB
	mdl *gorm.DB
}

func New(db *gorm.DB) *ManagerAccountRepository {
	return &ManagerAccountRepository{
		db:  db,
		mdl: db.Model(models.ManagerAccount{}).Session(&gorm.Session{}),
	}
}
