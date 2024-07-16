package port

import (
	"context"
	"go-hex-arch/internal/core/domain"
	"go-hex-arch/internal/core/dto"

	"github.com/google/uuid"
)

type ManagerAccountRepository interface {
	// GetManagerAccountList(query dtm.QueryContract, managerQuery dtm.QueryManagerAccount) (domain.ManagerAccountList, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error)
	GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error)
	Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error)

	// GetManagerAccountByEmail(email string) (domain.ManagerAccount, error)
	// GetDefaultManagerAccount() (domain.ManagerAccount, error)
	// CreateManagerAccount(manager domain.ManagerAccount) (domain.ManagerAccount, error)
	// UpdateManagerAccount(id string, manager domain.ManagerAccount) (domain.ManagerAccount, error)
	// DeleteManagerAccount(id string) error
}

type ManagerAccountService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error)
	GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error)
	Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error)
}
