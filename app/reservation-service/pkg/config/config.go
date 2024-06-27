package config

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/ztrue/tracerr"

	"bookit/pkg/database"
	"bookit/pkg/logger"
	"bookit/pkg/server"
)

var gist *Config

var flags = []string{"config", "envPrefix"}

type Config struct {
	App      App                  `mapstructure:"App"`
	Database database.PgEntConfig `mapstructure:"Psql"`
	Logger   logger.Config        `mapstructure:"Logger"`
	Server   server.FiberConfig   `mapstructure:"Server"`
}

type App struct {
	Mode        string `mapstructure:"Mode"`
	Version     string `mapstructure:"Version"`
	IsDebug     bool   `mapstructure:"IsDebug"`
	AuthEnabled bool   `mapstructure:"AuthEnabled"`
}

func Gist() *Config {
	if gist == nil {
		panic(tracerr.New("config not initialized"))
	}
	return gist
}
func New() *Config {
	gist = &Config{}
	conf := config.NewWithOptions("cfg", config.ParseEnv)
	conf.AddDriver(yaml.Driver)
	err := conf.LoadFlags(flags)
	if err != nil {
		panic(tracerr.Wrap(fmt.Errorf("failed to load flags: %w", err)))
	}
	err = conf.LoadFiles(conf.String("config"))
	if err != nil {
		panic(tracerr.Wrap(fmt.Errorf("failed to load config: %w", err)))
	}
	err = conf.Decode(&gist)
	if err != nil {
		panic(tracerr.Wrap(fmt.Errorf("failed to decode config: %w", err)))
	}
	return gist
}
