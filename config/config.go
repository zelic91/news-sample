package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	JWTKey      string `mapstructure:"JWT_KEY"`
	Port        int    `mapstructure:"PORT"`
	PostgresURL string `mapstructure:"POSTGRES_URL"`
}

func Init() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.BindEnv("JWT_KEY")
	viper.BindEnv("PORT")
	viper.BindEnv("POSTGRES_URL")

	viper.AutomaticEnv()

	// For .env file only
	err := viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		log.Printf("No .env file found. Using ENV instead.")
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
