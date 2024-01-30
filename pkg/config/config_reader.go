package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ConfigReader struct {
	configPath     string
	configFilename string
}

func NewConfigReader(configPath, configFilename string) *ConfigReader {
	return &ConfigReader{
		configPath:     configPath,
		configFilename: configFilename,
	}
}

func (cr *ConfigReader) ReadConfig(config any) error {
	log.Info("Reading data from config file")
	viper.AddConfigPath(cr.configPath)
	viper.SetConfigName(cr.configFilename)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read local config: %s", err)
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Failed to unmarshal data on local config: %s", err)
		return err
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return err
	}

	log.Infof("Local config successfully readed: %s", config)
	return nil
}
