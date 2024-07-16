package manager_account

import (
	"context"
	"go-hex-arch/internal/adapter/storage/mongo/models"
	"go-hex-arch/internal/core/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Create Manager ...
func (m *ManagerAccountRepository) Create(ctx context.Context, manager domain.ManagerAccount) (*domain.ManagerAccount, error) {
	var result domain.ManagerAccount
	var resultDbm models.ManagerAccount

	dbm := models.ManagerAccount{ID: uuid.New().String(), Name: manager.Name, Email: manager.Email, Phone: manager.Phone}

	insertResult, err := m.db.Collection("manager-account").InsertOne(ctx, dbm)
	if err != nil {
		return nil, err
	}

	err = m.db.Collection("manager-account").FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&resultDbm)
	if err != nil {
		return nil, err
	}

	objID, _ := uuid.Parse(resultDbm.ID)

	result = domain.ManagerAccount{ID: objID, Name: resultDbm.Name, Email: resultDbm.Email, Phone: resultDbm.Phone}

	return &result, nil
}
