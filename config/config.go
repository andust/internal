package config

import "github.com/kelseyhightower/envconfig"

type SystemConfig struct {
	Environment string
	LogLevel    string `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func InitConfig() (cfg SystemConfig, err error) {
	err = envconfig.Process("", &cfg)

	return
}
