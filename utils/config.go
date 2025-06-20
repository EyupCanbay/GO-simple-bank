package utils

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// the values are read by viper from a config file or environment variables
type Config struct {
	SBUri               string        `mapstructure:"DB_URI"`
	SBSoruce            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymetricKey    string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
