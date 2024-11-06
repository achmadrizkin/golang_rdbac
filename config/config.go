package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// MySQL Setup
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUsername string `mapstructure:"SQL_USER"`
	DBPassword string `mapstructure:"SQL_PASSWORD"`
	DBName     string `mapstructure:"SQL_DB"`
	DBPort     string `mapstructure:"SQL_PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
