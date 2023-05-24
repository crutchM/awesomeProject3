package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENV"`

	PgURL string `mapstructure:"POSTGRES_URL"`

	DefaultVideoPath string `mapstructure:"path"`

	Port string `mapstructure:"port"`
}

func NewConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigFile("/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err.Error())
	}

	return config

}
