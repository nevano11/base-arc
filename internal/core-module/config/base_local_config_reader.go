package config

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
)

type BaseLocalConfigReader struct {
}

func NewBaseLocalConfigReader() LocalConfigReader {
	return &BaseLocalConfigReader{}
}

func (cr *BaseLocalConfigReader) NewLocalConfig(configPath string) (*LocalConfig, error) {
	log.Info("Reading data from config file")
	localConfig := new(LocalConfig)

	configFile, err := os.Open(configPath)
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(configFile)

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)

	if err := jsonParser.Decode(&localConfig); err != nil {
		return nil, err
	}

	log.Info(localConfig)

	return localConfig, nil
}
