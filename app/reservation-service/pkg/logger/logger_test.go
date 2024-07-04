package logger

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateLogger(t *testing.T) {
	cfg := Config{
		Level: "debug",
	}
	logger := New(cfg)
	logger.Info().Msg("test")
	logger.Debug().Msg("test")
	logger.Error().Msg("test")
	logger = Gist(context.Background())
	logger.Warn().Msg("from gist")
}

func Test_ShouldGetLoggerFromCtx(t *testing.T) {
	cfg := Config{
		Level: "debug",
	}
	logger := New(cfg)
	sb := strings.Builder{}
	lg := logger.With().Str("method", "POST").Logger().Output(&sb)
	ctx := WithCtx(context.Background(), &lg)
	logger = Gist(ctx)
	logger.Info().Msg("test")
	if !assert.Contains(t, sb.String(), "POST") {
		t.Logf("logger output: %s", sb.String())
	}
}
