package manager_account

import (
	"context"
	"go-hex-arch/internal/adapter/storage/database/models"
	"go-hex-arch/internal/adapter/storage/database/serializers"
	"go-hex-arch/internal/core/domain"

	"github.com/google/uuid"
)

func (d *ManagerAccountRepository) Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error) {
	var manager models.ManagerAccount

	res := d.mdl.WithContext(ctx).First(&manager, id)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrDataNotFound
	}

	return serializers.ManagerAccountDomain(&manager), nil
}
