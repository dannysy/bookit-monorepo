package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"

	"bookit/pkg/errors"
)

type ctxLogger struct{}

var gist *zerolog.Logger

type Config struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func Gist(ctx context.Context) *zerolog.Logger {
	logger, ok := ctx.Value(ctxLogger{}).(*zerolog.Logger)
	if ok {
		return logger
	}
	if gist == nil {
		lg := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			With().Timestamp().Caller().Logger()
		lg.Warn().Msg("logger not configured!")
		return &lg
	}
	return gist
}

func WithCtx(ctx context.Context, lg *zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, lg)
}

func New(conf Config) *zerolog.Logger {
	level, err := zerolog.ParseLevel(conf.Level)
	if err != nil {
		panic(errors.Wrap(err, errors.WithMsg("failed to parse level")))
	}
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		Level(level).With().Timestamp().Caller().Logger()
	gist = &logger
	return gist
}
