package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type (
	Config struct {
		App   App   `mapstructure:"app"`
		Redis Redis `mapstructure:"redis"`
	}

	App struct {
		Env         string        `mapstructure:"env"`
		HTTPPort    int           `mapstructure:"http_port"`
		HTTPTimeout time.Duration `mapstructure:"http_timeout"`
	}

	Postgres struct {
		DbHost            string `mapstructure:"db_host"`
		DbPort            string `mapstructure:"db_port"`
		DbUser            string `mapstructure:"db_user"`
		DbPass            string `mapstructure:"db_pass"`
		DbName            string `mapstructure:"db_name"`
		DbSchema          string `mapstructure:"db_schema"`
		MaxOpenConnection int    `mapstructure:"maxOpenConnections"`
		MaxIdleConnection int    `mapstructure:"maxIdleConnections"`
	}

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	}
)

// New Init load config from config file
func New() (cfg Config, err error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return cfg, nil
}
