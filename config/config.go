package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	MongodbURL string
	Database   string
}

func LoadConfig(filename string) (Config, error) {
	// Set the config file name and path
	viper.SetConfigFile(filename)

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse the config into a struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
