package crdb

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (c *Config) SetupSchema(ctx context.Context) error {
	log.Info().Str(`dbname`, c.DbName).Msg(`recreating database`)
	err := c.DropDatabase(ctx, c.DbName)
	if err != nil {
		return err
	}
	err = c.CreateDatabase(ctx, c.DbName)
	if err != nil {
		return err
	}
	err = c.Close()
	if err != nil {
		return err
	}
	err = c.Init()
	if err != nil {
		return err
	}
	defer c.Close()
	_, err = c.db.ExecContext(ctx, createUsersStmt)
	if err != nil {
		return err
	}
	return nil
}