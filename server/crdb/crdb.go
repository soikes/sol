package crdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = `localhost`
	port = `26257`
	user = `root`
)

var (
	ErrNotInitialized = errors.New(`db has not been initialized`)
)

type Config struct {
	db *sql.DB
	initialized bool
}

func (c *Config) Init() error {
	return c.initWithDb(``)
}

func (c *Config) InitWith(name string) error {
	return c.initWithDb(name)
}

func (c *Config) initWithDb(name string) error {
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s sslmode=disable`,
		host, port, user)
	if name != `` {
		psqlInfo = psqlInfo + fmt.Sprintf(` dbname=%s`, name)
	}
	db, err := sql.Open(`postgres`, psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	c.db = db
	c.initialized = true
	return nil
}

func (c *Config) Close() error {
	return c.db.Close()
}

func (c *Config) UseDatabase(ctx context.Context, name string) error {
	if !c.initialized {
		return ErrNotInitialized
	}
	_, err := c.db.ExecContext(ctx, fmt.Sprint(`USE DATABASE `), name)
	return err
}

func (c *Config) CreateDatabase(ctx context.Context, name string) error {
	_, err := c.db.ExecContext(ctx, fmt.Sprint(`CREATE DATABASE IF NOT EXISTS `, name))
	return err
}

func (c *Config) CreateTable(ctx context.Context, name string) error {
	_, err := c.db.ExecContext(ctx, fmt.Sprint(`CREATE TABLE IF NOT EXISTS `, name))
	return err
}

func (c *Config) CreateUser(ctx context.Context, name string) error {
	_, err := c.db.ExecContext(ctx, `CREATE USER IF NOT EXISTS ?`, name)
	return err
}

func (c *Config) DropDatabase(ctx context.Context, name string) error {
	_, err := c.db.ExecContext(ctx, fmt.Sprint(`DROP DATABASE IF EXISTS `, name))
	return err
}

func (c *Config) ExecContext(ctx context.Context, stmt string) error {
	_, err := c.db.ExecContext(ctx, stmt)
	return err
}