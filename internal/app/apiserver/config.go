package apiserver

import "github.com/IFtech-A/http_rest_api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string        `yaml:"bind_addr"`
	LogLevel string        `yaml:"log_level"`
	Store    *store.Config `yaml:"store"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
