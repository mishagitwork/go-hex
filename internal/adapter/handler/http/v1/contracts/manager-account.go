package contracts

import "github.com/google/uuid"

// ManagerAccountCreateContract ...
type ManagerAccountContract struct {
	Image         string `json:"image"`
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	MessengerLink string `json:"messengerLink"`
	Phone         string `json:"phone" binding:"required"`
}

// ManagerAccountCreateContract ...
type ManagerAccountCreateAndUpdateContract struct {
	ManagerAccountContract
	Password string `json:"password" binding:"required"`
}

// ManagerAccountResponseContract ...
type ManagerAccountResponseContract struct {
	ID uuid.UUID `json:"id"`
	ManagerAccountContract
}
