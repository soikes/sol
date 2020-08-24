package main

import (
	"context"

	"soikke.li/sol/crdb"
)

func cmdSetup(ctx context.Context) error {
	cfg := crdb.Config{}
	return cfg.SetupSchema(ctx)
}
