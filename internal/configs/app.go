package configs

type App struct {
	DebugMode bool `env-required:"true" env:"DEBUG_MODE"`
	HTTPPort  uint `env-required:"true" env:"HTTP_PORT"`
	GRPCPort  uint `env-required:"true" env:"GRPC_PORT"`
}
