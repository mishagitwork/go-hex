package serializers

import (
	"go-hex-arch/internal/adapter/storage/database/models"
	"go-hex-arch/internal/core/domain"
)

func ManagerAccountModel(d *domain.ManagerAccount) *models.ManagerAccount {
	return &models.ManagerAccount{
		ID:              d.ID,
		Name:            d.Name,
		Image:           d.Image,
		Phone:           d.Phone,
		Email:           d.Email,
		Password:        d.Password,
		MessengerLink:   d.MessengerLink,
		AdditionalPhone: d.AdditionalPhone,
	}
}

func ManagerAccountDomain(d *models.ManagerAccount) *domain.ManagerAccount {
	return &domain.ManagerAccount{
		ID:              d.ID,
		Name:            d.Name,
		Image:           d.Image,
		Phone:           d.Phone,
		Email:           d.Email,
		Password:        d.Password,
		MessengerLink:   d.MessengerLink,
		AdditionalPhone: d.AdditionalPhone,
	}
}

func ManagerAccountDomains(d []models.ManagerAccount) []domain.ManagerAccount {
	res := make([]domain.ManagerAccount, 0, len(d))
	for _, s := range d {
		res = append(res, *ManagerAccountDomain(&s))
	}

	return res
}
