package main

import (
	"bookit/cmd"
	"bookit/pkg/config"
	"bookit/pkg/logger"

	_ "github.com/lib/pq"
)

// main команда запуска go run main.go -config ./etc/config.yml
//
//	@title						Reservation REST API
//	@version					1.0
//
//	@host						localhost:8084
//	@schemes					http
//
//	@BasePath					/v1
//
//	@securityDefinitions.apikey	AccessToken
//	@in							header
//	@name						authorization
func main() {
	cfg := config.New()
	log := logger.New(cfg.Logger)
	cmds := map[string]func(config *config.Config){
		"api": cmd.NewApi,
	}
	c, ok := cmds[cfg.App.Mode]
	if !ok {
		log.Fatal().Str("app.mode", cfg.App.Mode).Msg("unknown mode")
	}
	log.Info().Str("app.mode", cfg.App.Mode).Msg("starting app")
	c(cfg)
	log.Info().Str("app.mode", cfg.App.Mode).Msg("app finished")
}
