package manager_account

import (
	"context"
	"go-hex-arch/internal/adapter/storage/database/models"
	"go-hex-arch/internal/adapter/storage/database/serializers"
	"go-hex-arch/internal/adapter/storage/database/utils"
	"go-hex-arch/internal/core/domain"
	"go-hex-arch/internal/core/dto"
	"strings"
)

// GetManagerAccountList ...
func (d *ManagerAccountRepository) GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error) {

	var list dto.List[domain.ManagerAccount]
	var data []models.ManagerAccount
	var count int64

	db := d.mdl.WithContext(ctx)

	if query.Search != "" {
		db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query.Search)+"%").
			Or("LOWER(email) LIKE ?", "%"+strings.ToLower(query.Search)+"%")
	}

	if managerQuery.Roles != "" {
		db.Where("role in (?)", strings.Split(managerQuery.Roles, ","))
	}

	if len(managerQuery.AdditionalRoles) > 0 {
		db.Where("additional_roles && ?", managerQuery.AdditionalRoles)
	}

	if managerQuery.AvailableUponRegistration != "" {
		db = db.Where("available_upon_registration = ?", managerQuery.AvailableUponRegistration)
	}

	if len(managerQuery.ProjectPermissions) > 0 {
		db = db.Where("string_to_array(project_permission, ',') && ?", managerQuery.ProjectPermissions)
	}

	res := db.Count(&count)
	if res.Error != nil {
		return nil, res.Error
	}

	res = db.Scopes(utils.Paginate(query), utils.OrderBy(query)).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}

	list.Data = serializers.ManagerAccountDomains(data)
	list.Meta.Count = count

	return &list, nil
}
