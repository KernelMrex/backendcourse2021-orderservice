package main

import (
	"github.com/kelseyhightower/envconfig"
)

const appID = "orderservice"

type config struct {
	ServeRESTAddress   string `envconfig:"serve_rest_address" default:":8000"`
	DataSourceHostname string `envconfig:"data_source_hostname" default:""`
	DataSourceUsername string `envconfig:"data_source_username" default:"root"`
	DataSourcePassword string `envconfig:"data_source_password" default:"1234"`
	DataSourceDatabase string `envconfig:"data_source_database" default:"orderservice"`
}

func parseEnv() (*config, error) {
	cfg := new(config)
	if err := envconfig.Process(appID, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
