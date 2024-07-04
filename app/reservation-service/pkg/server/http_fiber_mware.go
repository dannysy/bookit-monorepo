package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/google/uuid"

	"bookit/pkg/logger"
)

func Recover() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c fiber.Ctx, e any) {
			err := fmt.Errorf("%v", e)
			logger.Gist(c.UserContext()).Err(err).Msg("global recover catch panic")
		},
	})
}

func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		traceID := c.Get("X-Trace-Id", uuid.NewString())
		ctx := c.UserContext()
		lg := logger.Gist(ctx).With().
			Str("trace-id", traceID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Logger()
		ctx = logger.WithCtx(ctx, &lg)
		c.SetUserContext(ctx)
		c.Set("X-Trace-Id", traceID)
		start := time.Now()
		statusCode := new(int)
		defer func() {
			lg.Debug().Str("duration", time.Since(start).String()).Int("status", *statusCode).Msg("call")
		}()
		err := c.Next()
		// TODO: implement middleware error handler
		//if err = c.App().ErrorHandler(c, err); err != nil {
		//	return c.SendStatus(fiber.StatusInternalServerError)
		//}
		*statusCode = c.Response().StatusCode()
		return err
	}
}
