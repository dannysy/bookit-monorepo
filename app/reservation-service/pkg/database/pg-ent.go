package database

import (
	"context"

	"github.com/ztrue/tracerr"

	"bookit/internal/repo/ent"
	"bookit/internal/repo/ent/migrate"
	"bookit/pkg/logger"
)

type PgEntConfig struct {
	URL          string `mapstructure:"URL"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
}

type PgEnt struct {
	*ent.Client
}

func NewPgEnt(cfg PgEntConfig) *PgEnt {
	ctx := context.Background()
	client, err := ent.Open("postgres", cfg.URL)
	if err != nil {
		logger.Gist(ctx).Fatal().Err(tracerr.Wrap(err)).Msg("failed opening connection to postgres")
		return nil
	}
	if err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		logger.Gist(ctx).Fatal().Err(tracerr.Wrap(err)).Msg("failed creating schema resources")
	}

	return &PgEnt{
		Client: client,
	}
}

func (p *PgEnt) Close() error {
	return p.Client.Close()
}
