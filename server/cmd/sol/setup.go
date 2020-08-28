package main

import (
	"context"

	"soikke.li/sol"
)

func cmdSetup(ctx context.Context, cfg *sol.Config) error {
	return cfg.Crdb.SetupSchema(ctx)
}
