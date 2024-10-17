package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   *Server   `mapstructure:"server" validate:"required"`
	Database *Database `mapstructure:"database" validate:"required"`
}

type Server struct {
	Port         int      `mapstructure:"port" validate:"required"`
	AllowOrigins []string `mapstructure:"allowOrigins" validate:"required"`
	BodyLimit    string   `mapstructure:"bodyLimit" validate:"required"`
	Timeout      int      `mapstructure:"timeout" validate:"required"`
}

type Database struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DBName   string `mapstructure:"dbname" validate:"required"`
	SSLMode  string `mapstructure:"sslmode" validate:"required"`
	Schema   string `mapstructure:"schema" validate:"required"`
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}
