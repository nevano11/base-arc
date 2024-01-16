package config

import "github.com/labstack/gommon/log"

type LocalConfigReader interface {
	NewLocalConfig(configPath string) (*LocalConfig, error)
}

type FakeLocalConfigReader struct {
}

func NewFakeLocalConfigReader() LocalConfigReader {
	return &FakeLocalConfigReader{}
}

func (cr *FakeLocalConfigReader) NewLocalConfig(configPath string) (*LocalConfig, error) {
	log.Info("Reading data from config file. Unmarshall it on LocalConfig struct")
	return &LocalConfig{
		GrpcAddr: "localhost:9000",
		TarantoolConfig: TarantoolConfig{
			Address:  "localhost:9000",
			User:     "name",
			Password: "pass",
		},
	}, nil
}
