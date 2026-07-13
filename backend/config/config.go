package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

var AppConfig *Config

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	AppConfig = &cfg

	return nil
}

func GetConfig() *Config {
	return AppConfig
}
