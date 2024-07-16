package service

import (
	"context"
	"errors"
	"go-hex-arch/internal/core/domain"
	"go-hex-arch/internal/core/dto"
	"go-hex-arch/internal/core/port"

	"github.com/google/uuid"
)

type ManagerAccountMongoService struct {
	repo                  port.ManagerAccountRepository
	managerAccountService port.ManagerAccountService
}

func NewManagerAccountMongoService(repo port.ManagerAccountRepository, managerAccountService port.ManagerAccountService) *ManagerAccountMongoService {
	return &ManagerAccountMongoService{
		repo,
		managerAccountService,
	}
}

// GetManagerAccount ...
func (s *ManagerAccountMongoService) Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error) {
	if res, err := s.managerAccountService.Get(ctx, id); err == nil {
		return res, nil
	}

	return s.repo.Get(ctx, id)
}

func (s *ManagerAccountMongoService) GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error) {
	return s.repo.GetList(ctx, query, managerQuery)
}

// CreateManagerAccount ...
func (s *ManagerAccountMongoService) Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error) {
	return s.repo.Create(ctx, manager)
}

// TestFunctionExampleDivide ...
func (s *ManagerAccountMongoService) FunctionExampleDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("zero dividing")
	}
	return a / b, nil
}
