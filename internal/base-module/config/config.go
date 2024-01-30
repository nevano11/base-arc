package config

import "fmt"

type Config struct {
	GrpcAddr string `mapstructure:"core-addr" validate:"required"`
}

func (lc *Config) String() string {
	return fmt.Sprintf("Config {GrpcAddr=%s}", lc.GrpcAddr)
}
