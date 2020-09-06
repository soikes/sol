package crdb

import (
	"context"
)

func (cfg *Config) CreateUser(ctx context.Context, name, email, password string) error {
	err := cfg.ExecContext(ctx, insertUserStmt, name, email, password)
	return err
}

func (cfg *Config) GetUserPassword(ctx context.Context, email string) (string, error) {
	rows, err := cfg.QueryContext(ctx, getUserPasswordStmt, email)
	if err != nil {
		return ``, err
	}
	var pw string
	ok := rows.Next()
	if !ok {
		return ``, ErrNoRows
	}
	err = rows.Scan(&pw)
	if err != nil {
		return ``, err
	}
	return pw, nil
}