package crdb

import (
	"context"
)

func (cfg *Config) SetupSchema(ctx context.Context) error {
	cfg.Log.Info().Str(`dbname`, cfg.DbName).Msg(`recreating database`)
	err := cfg.DropDatabase(ctx, cfg.DbName)
	if err != nil {
		return err
	}
	err = cfg.CreateDatabase(ctx, cfg.DbName)
	if err != nil {
		return err
	}
	err = cfg.Close()
	if err != nil {
		return err
	}
	err = cfg.InitDB()
	if err != nil {
		return err
	}
	defer cfg.Close()
	_, err = cfg.db.ExecContext(ctx, createUsersTableStmt)
	if err != nil {
		return err
	}
	return nil
}