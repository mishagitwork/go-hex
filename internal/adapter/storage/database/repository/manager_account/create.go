package manager_account

import (
	"context"
	"go-hex-arch/internal/adapter/storage/database/serializers"
	"go-hex-arch/internal/core/domain"
)

// CreateManagerAccount ...
func (d *ManagerAccountRepository) Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error) {
	model := serializers.ManagerAccountModel(&manager)

	res := d.mdl.WithContext(ctx).Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, domain.ErrNoCreatedData
	}

	return serializers.ManagerAccountDomain(model), nil
}
