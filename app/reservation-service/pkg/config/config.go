package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"

	"bookit/pkg/database"
	"bookit/pkg/errors"
	"bookit/pkg/iam"
	"bookit/pkg/logger"
	"bookit/pkg/server"
)

var gist *Config

var flags = []string{"config", "envPrefix"}

type Config struct {
	App      App                  `mapstructure:"App"`
	Iam      iam.CasdoorConfig    `mapstructure:"Casdoor"`
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
		panic(errors.New("config not initialized"))
	}
	return gist
}
func New() *Config {
	gist = &Config{}
	conf := config.NewWithOptions("cfg", config.ParseEnv, config.ParseTime)
	conf.AddDriver(yaml.Driver)
	err := conf.LoadFlags(flags)
	if err != nil {
		panic(errors.Wrap(err, errors.WithMsg("failed to load flags")))
	}
	err = conf.LoadFiles(conf.String("config"))
	if err != nil {
		panic(errors.Wrap(err, errors.WithMsg("failed to load config")))
	}
	err = conf.Decode(&gist)
	if err != nil {
		panic(errors.Wrap(err, errors.WithMsg("failed to decode config")))
	}
	return gist
}
