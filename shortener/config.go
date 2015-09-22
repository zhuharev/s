package shortener

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Web struct {
		Http     bool
		Port     int
		Endpoint string
	}
	Database struct {
		DataDir string
	}
}

func NewConfig(filepath string) (*Config, error) {
	cfg := new(Config)
	e := gcfg.ReadFileInto(cfg, filepath)
	return cfg, e
}
