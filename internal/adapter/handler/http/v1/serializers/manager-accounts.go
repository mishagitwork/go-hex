package serializers

import (
	"go-hex-arch/internal/adapter/handler/http/v1/contracts"
	"go-hex-arch/internal/core/domain"
)

func ManagerAccountResponse(d *domain.ManagerAccount) *contracts.ManagerAccountResponseContract {

	return &contracts.ManagerAccountResponseContract{
		ID: d.ID,
		ManagerAccountContract: contracts.ManagerAccountContract{
			Image:         d.Image,
			Name:          d.Name,
			Phone:         d.Phone,
			Email:         d.Email,
			MessengerLink: d.MessengerLink,
		},
	}
}
