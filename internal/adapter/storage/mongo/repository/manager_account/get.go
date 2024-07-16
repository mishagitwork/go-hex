package manager_account

import (
	"context"
	"go-hex-arch/internal/core/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Get ...
func (m *ManagerAccountRepository) Get(ctx context.Context, id uuid.UUID) (*domain.ManagerAccount, error) {
	var result domain.ManagerAccount

	filter := bson.M{"_id": id.String()}

	err := m.db.Collection("manager-account").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
