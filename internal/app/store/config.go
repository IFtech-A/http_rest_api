package store

// Config ...
type Config struct {
	DatabaseURL string `yaml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
