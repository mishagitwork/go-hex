package manager_account

import (
	"context"
	"go-hex-arch/internal/adapter/storage/mongo/repository/utils"
	"go-hex-arch/internal/core/domain"
	"go-hex-arch/internal/core/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get Manager List ...
func (m *ManagerAccountRepository) GetList(ctx context.Context, query dto.Query, managerQuery dto.QueryManagerAccount) (*dto.List[domain.ManagerAccount], error) {
	var result dto.List[domain.ManagerAccount]
	var data []domain.ManagerAccount

	opt := options.Find()
	countOpt := options.Count().SetHint("_id_")

	utils.SetPagination(opt, query)
	utils.SetSort(opt, query)

	filter := bson.M{}

	count, err := m.db.Collection("manager-account").CountDocuments(ctx, filter, countOpt)
	if err != nil {
		return nil, err
	}

	cur, err := m.db.Collection("manager-account").Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	result.Meta.Count = count
	result.Data = data

	return &result, nil
}
