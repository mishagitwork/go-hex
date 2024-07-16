package database

import (
	"go-hex-arch/internal/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	*gorm.DB
}

// New ...
func New(config *configs.Database) (*gorm.DB, error) {

	dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// if config.AutoMigrate {
	// 	err = migrations.ApplyMigrations(db)
	// 	if err != nil {
	// 		log.Fatal().Msg(err.Error())
	// 	}
	// }

	return db, nil
}
