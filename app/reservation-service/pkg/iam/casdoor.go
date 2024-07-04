package iam

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/ztrue/tracerr"
)

var gist *casdoorsdk.Client

type CasdoorConfig struct {
	ClientId     string `mapstructure:"ClientId"`
	ClientSecret string `mapstructure:"ClientSecret"`
	Cert         string `mapstructure:"Cert"`
	Url          string `mapstructure:"Url"`
	Organization string `mapstructure:"Organization"`
	Application  string `mapstructure:"Application"`
}

func NewCasdoor(cfg CasdoorConfig) *casdoorsdk.Client {
	gist = casdoorsdk.NewClient(cfg.Url, cfg.ClientId, cfg.ClientSecret, cfg.Cert, cfg.Organization, cfg.Application)
	return gist
}

func Gist() *casdoorsdk.Client {
	if gist == nil {
		panic(tracerr.New("casdoor client not initialized"))
	}
	return gist
}
