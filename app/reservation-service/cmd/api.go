package cmd

import (
	"bookit/pkg/app"
	"bookit/pkg/config"
	"bookit/pkg/database"
	"bookit/pkg/server"
)

type ApiCommand struct {
	cmd *app.Command
}

func NewApi(config *config.Config) {
	cmd := app.NewCommand(
		[]app.Closable{database.NewPgEnt(config.Database)},
		[]app.Runnable{server.NewFiberServer(config.Server)},
	)
	cmd.Do()
}
