package main

import (
	"context"

	"soikke.li/sol/config"
)

func cmdSetup(ctx context.Context, cfg *config.Config) error {
	return cfg.Crdb.SetupSchema(ctx)
}
