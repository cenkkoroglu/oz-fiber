package config

import (
	"github.com/cenkkoroglu/oz-fiber/app/models"
	"github.com/cenkkoroglu/oz-fiber/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config *models.Config

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Fatal error when reading config file", zap.Error(err))
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.Error("Fatal error when decoding config file into struct", zap.Error(err))
		return err
	}

	return nil
}

func GetConfig() *models.Config {
	return Config
}
