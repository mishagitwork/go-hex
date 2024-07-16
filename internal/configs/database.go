package configs

type Database struct {
	Host        string `env-required:"true" env:"DB_HOST"`
	Port        uint   `env-required:"true" env:"DB_PORT"`
	Username    string `env-required:"true" env:"DB_USERNAME"`
	Password    string `env-required:"true" env:"DB_PASSWORD"`
	Database    string `env-required:"true" env:"DB_DATABASE"`
	AutoMigrate bool   `env-required:"true" env:"DB_AUTOMIGRATE"`
}
