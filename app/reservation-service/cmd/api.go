package cmd

import (
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/utils/v2"

	"bookit/internal/transport"
	"bookit/pkg/app"
	"bookit/pkg/config"
	"bookit/pkg/database"
	"bookit/pkg/iam"
	"bookit/pkg/server"
)

type ApiCommand struct {
	cmd *app.Command
}

func NewApi(config *config.Config) {
	// init http server & routes
	srv := server.NewFiberServer(config.Server)
	srv.Use(
		server.Logger(),
		server.Error(),
		transport.Auth(),
		cors.New(
			cors.Config{
				AllowOrigins: "*",
				AllowHeaders: "X-Trace-Id, Content-Type, Authorization, Origin, Accept",
				AllowMethods: "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS",
			}),
		server.Recover(),
		requestid.New(requestid.Config{
			Header:    "X-Request-Id",
			Generator: utils.UUIDv4,
		}),
	)
	// init routes
	router := srv.Group("/v1")
	router.Get("/version", transport.Version)
	router.Get("/healthz", transport.Health)
	router.Get("/panic", transport.Panic)
	router.Get("/error", transport.Error)
	router.Get("/signin", transport.Signin)
	router.Get("/user", transport.User)
	_ = iam.NewCasdoor(config.Iam)
	// start application
	cmd := app.NewCommand(
		[]app.Closable{database.NewPgEnt(config.Database)},
		[]app.Runnable{srv},
	)
	cmd.Do()
}
