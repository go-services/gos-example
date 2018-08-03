package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/go-services/core"
	"github.com/kevinburke/go.uuid"
)

type ServiceConfig struct {
	core.DefaultServiceConfig
	DBConnectionString string `envconfig:"db_connection_string" default:":8081"`
}

var cf *ServiceConfig

func Get() *ServiceConfig {
	if cf == nil {
		cf = &ServiceConfig{}
		envconfig.MustProcess("user", cf)
		if cf.Name == "" {
			cf.Name = "user"
		}
		if cf.Instance == "" {
			cf.Instance = uuid.NewV4().String()
		}
	}
	return cf
}
