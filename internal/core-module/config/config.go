package config

import "fmt"

type TarantoolConfig struct {
	Address  string `mapstructure:"addr" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

func (lc *TarantoolConfig) String() string {
	return fmt.Sprintf("TarantoolConfig {Address=%s, User=%s, Password=%s}", lc.Address, lc.User, lc.Password)
}

type Config struct {
	GrpcAddr        string          `mapstructure:"core-addr" validate:"required"`
	TarantoolConfig TarantoolConfig `mapstructure:"tarantool" validate:"required"`
}

func (lc *Config) String() string {
	return fmt.Sprintf("Config {GrpcAddr=%s, TarantoolConfig=%s}", lc.GrpcAddr, lc.TarantoolConfig.String())
}
