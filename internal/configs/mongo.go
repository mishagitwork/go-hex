package configs

type Mongo struct {
	Host     string `env-required:"true" env:"MONGO_HOST"`
	Port     uint   `env-required:"true" env:"MONGO_PORT"`
	Username string `env-required:"true" env:"MONGO_USERNAME"`
	Password string `env-required:"true" env:"MONGO_PASSWORD"`
	Database string `env-required:"true" env:"MONGO_DATABASE"`
}
