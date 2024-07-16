package models

// ManagerAccount ...
type ManagerAccount struct {
	ID            string `bson:"_id"`
	Image         string
	Name          string
	Email         string
	Password      string
	MessengerLink string
	Phone         string
}
