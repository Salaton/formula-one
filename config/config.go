package config

import (
	"fmt"

	"github.com/Salaton/formula-one/pkg/presentation/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	configFile = ".env"
)

type Configuration struct {
	EnvConfig EnvVariablesConfig `mapstructure:",squash"`
}

type EnvVariablesConfig struct {
	ErgastAPIResponseType string `mapstructure:"ERGAST_API_RESPONSE_TYPE"`
	ErgastAPIEndpoint     string `mapstructure:"ERGAST_API_ENDPOINT"`
}

func LoadConfig() (*Configuration, error) {
	var config *Configuration
	logger := logger.NewLogger()
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	// Watch config file, no need to restart server after updating the config file
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Log(zerolog.InfoLevel, "Config file changed: %v", e.Name)
	})
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found, please update the `ConfigPath`: %w", err)
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the config: %w", err)
	}

	return config, nil
}