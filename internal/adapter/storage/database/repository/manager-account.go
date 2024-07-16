package repository

import (
	"go-hex-arch/internal/adapter/storage/database/repository/manager_account"

	"gorm.io/gorm"
)

func NewManagerAccountRepository(db *gorm.DB) *manager_account.ManagerAccountRepository {
	return manager_account.New(db)
}
