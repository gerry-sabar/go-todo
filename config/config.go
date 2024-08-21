package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	APPPort    string `mapstructure:"APP_PORT"`
	APPEnv     string `mapstructure:"APP_ENV"`
}

var Cfg *Config

func LoadConfig() {
	Cfg = &Config{}
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.SetConfigType("env")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("reading environment variables, failed to read config from file: %v", err))
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}

}
