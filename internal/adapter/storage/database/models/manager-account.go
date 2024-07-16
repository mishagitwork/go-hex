package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// ManagerAccount ...
type ManagerAccount struct {
	gorm.Model `json:"-" `

	ID                        uuid.UUID      `gorm:"primaryKey;type:uuid" `
	Image                     string         `gorm:"type:varchar" `
	Name                      string         `gorm:"type:varchar;not null" `
	Email                     string         `gorm:"type:varchar;unique;not null" `
	Password                  string         `gorm:"not null" `
	MessengerLink             string         `gorm:"type:varchar" `
	Phone                     string         `gorm:"type:varchar;not null" `
	AdditionalPhone           string         `gorm:"type:varchar" `
	ProjectPermission         string         `gorm:"type:varchar;not null;default:'';"`
	AmoCRMID                  int64          `gorm:"type:integer;" `
	ManagerTag                string         `gorm:"type:varchar;not null;default:'';" `
	ManagerTagEnabled         bool           `gorm:"type:boolean;not null;default:false"`
	AvailableUponRegistration bool           `gorm:"type:boolean;not null;default:false" `
	MoySkladID                string         `gorm:"type:varchar;not null;default:''"`
	Role                      string         `gorm:"type:varchar;not null;default:'MANAGER'" `
	AdditionalRoles           pq.StringArray `gorm:"type:varchar[];not null;default: array[]::varchar[]"`
}

func (manager *ManagerAccount) BeforeSave(tx *gorm.DB) (err error) {

	fmt.Println("Before Save")

	return
}

func (manager *ManagerAccount) BeforeCreate(tx *gorm.DB) (err error) {

	manager.ID = uuid.New()

	return
}

func (manager *ManagerAccount) AfterFind(tx *gorm.DB) (err error) {

	if manager.AmoCRMID == 0 {
		manager.AmoCRMID = 1469293
	}

	return
}
