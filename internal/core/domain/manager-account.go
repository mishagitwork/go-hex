package domain

import (
	"github.com/google/uuid"
)

// ManagerAccount ...
type ManagerAccount struct {
	ID              uuid.UUID
	Image           string
	Name            string
	Email           string
	Password        string
	MessengerLink   string
	Phone           string
	AdditionalPhone string
}
