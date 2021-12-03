package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port         int `env:"PORT" envDefault:"8080"`
	PortInternal int `env:"PORT_INTERNAL" envDefault:"8585"`

	HDFS struct {
		Host string `env:"HDFS_HOST"`
		Port int    `env:"HDFS_PORT"`
	}
}

func New() (*Config, error) {
	c := new(Config)
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return c, nil
}