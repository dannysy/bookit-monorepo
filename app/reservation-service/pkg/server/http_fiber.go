package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

type FiberConfig struct {
	Host         string        `mapstructure:"Host"`
	Port         int           `mapstructure:"Port"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

type FiberServer struct {
	cfg FiberConfig
	*fiber.App
}

func NewFiberServer(config FiberConfig) *FiberServer {
	return &FiberServer{
		cfg: config,
		App: fiber.New(fiber.Config{
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		}),
	}
}

func (s *FiberServer) Stop() error {
	return s.App.Shutdown()
}

func (s *FiberServer) Start() error {
	return s.App.Listen(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}
