package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBURL      string `mapstructure:"DB_URL"`
	SigningKey string `json:"SIGNING_KEY"`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return
	}

	var config Config

	err = viper.Unmarshal(&config)

	if err != nil {
		return
	}

	if config.DBURL != "" {
		viper.Set("DB_URL", config.DBURL)
	}

	if config.SigningKey != "" {
		viper.Set("SIGNING_KEY", config.SigningKey)
	}
	return
}
