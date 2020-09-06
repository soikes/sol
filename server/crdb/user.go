package crdb

import (
	"context"

	"soikke.li/sol/svc/web"
)

func (cfg *Config) CreateUser(ctx context.Context, name, email, password string) (string, error) {
	rows, err := cfg.QueryContext(ctx, insertUserStmt, name, email, password)
	if err != nil {
		return ``, err
	}
	var id string
	ok := rows.Next()
	if !ok {
		return ``, ErrNoRows
	}
	err = rows.Scan(&id)
	if err != nil {
		return ``, err
	}
	return id, nil
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

func (cfg *Config) GetUserInfo(ctx context.Context, id string, u *web.User) error {
	rows, err := cfg.QueryContext(ctx, getUserInfoStmt, id)
	if err != nil {
		return err
	}
	ok := rows.Next()
	if !ok {
		return ErrNoRows
	}
	return rows.StructScan(u)
}