package crdb

import (
	"context"
)

func (c *Config) SetupSchema(ctx context.Context) error {
	err := c.Init()
	if err != nil {
		return err
	}
	err = c.DropDatabase(ctx, `sol`)
	if err != nil {
		return err
	}
	err = c.CreateDatabase(ctx, `sol`)
	if err != nil {
		return err
	}
	err = c.Close()
	if err != nil {
		return err
	}
	err = c.InitWith(`sol`)
	if err != nil {
		return err
	}
	defer c.Close()
	err = c.ExecContext(ctx, createUsersStmt)
	if err != nil {
		return err
	}
	return nil
}