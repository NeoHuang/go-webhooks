package server

type Config struct {
	Port int
}

func NewConfig(configFile string) *Config {
	return DefaultConfig()
}

func DefaultConfig() *Config {
	return &Config{
		Port: 8401,
	}
}
