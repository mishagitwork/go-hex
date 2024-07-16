package service

import (
	"context"
	"go-hex-arch/internal/core/domain"
	"go-hex-arch/internal/core/dto"
	"go-hex-arch/internal/core/port"

	"github.com/google/uuid"
)

type ManagerAccountService struct {
	db port.ManagerAccountRepository
}

func NewCategoryService(db port.ManagerAccountRepository) *ManagerAccountService {
	return &ManagerAccountService{
		db,
	}
}

// GetManagerAccount ...
func (s *ManagerAccountService) Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error) {
	return s.db.Get(ctx, id)
}

func (s *ManagerAccountService) GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error) {
	return s.db.GetList(ctx, query, managerQuery)
}

// CreateManagerAccount ...
func (s *ManagerAccountService) Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error) {
	return s.db.Create(ctx, manager)
}
