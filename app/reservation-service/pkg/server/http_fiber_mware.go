package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/google/uuid"

	apperrors "bookit/pkg/errors"
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

// Logger middleware
// This middleware must be registered on the top for correct duration calculation
// and before Error middleware for correct http status logging
func Logger() fiber.Handler {
	return func(c fiber.Ctx) error {
		traceID := c.Get("X-Trace-Id", uuid.NewString())
		ctx := c.UserContext()
		lg := logger.Gist(ctx).With().
			Str("trace-id", traceID).
			Logger()
		ctx = logger.WithCtx(ctx, &lg)
		c.SetUserContext(ctx)
		c.Set("X-Trace-Id", traceID)
		start := time.Now()
		defer func() {
			lg.Debug().
				Str("method", c.Method()).
				Str("path", c.Path()).
				Str("duration", time.Since(start).String()).
				Int("status", c.Response().StatusCode()).
				Msg("call")
		}()
		return c.Next()
	}
}

// Error middleware
// This middleware must be registered on the top but before Logger middleware
func Error() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}
		var appErr apperrors.Error
		if errors.As(err, &appErr) {
			logger.Gist(c.UserContext()).Error().
				Msg(appErr.Msg())
			logger.Gist(c.UserContext()).Printf("stacktrace:\n%s", appErr.StackTrace())
			return c.Status(appErr.HttpStatus()).JSON(appErr.Msg())
		}
		logger.Gist(c.UserContext()).Warn().Msg("err is not bookit/pkg/errors type!")
		logger.Gist(c.UserContext()).Error().Err(err).Send()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}
