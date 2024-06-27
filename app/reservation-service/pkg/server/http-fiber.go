package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type FiberConfig struct {
	Host string `mapstructure:"Host"`
	Port int    `mapstructure:"Port"`
}

type FiberServer struct {
	cfg FiberConfig
	*fiber.App
}

func NewFiberServer(config FiberConfig) *FiberServer {
	return &FiberServer{
		cfg: config,
		App: fiber.New(),
	}
}

func (s *FiberServer) Stop() error {
	return s.App.Shutdown()
}

func (s *FiberServer) Start() error {
	return s.App.Listen(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}
