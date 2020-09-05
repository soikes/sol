package crdb

import (
	"context"
)

func (cfg *Config) CreateUser(ctx context.Context, name, email string) error {
	_, err := cfg.db.ExecContext(ctx, insertUserStmt, name, email)
	return err
}