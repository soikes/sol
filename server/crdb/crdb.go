package crdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	ErrNotInitialized = errors.New(`db has not been initialized`)
	ErrNoRows         = errors.New(`no rows in result set`)
)

type Config struct {
	Host string
	Port string
	User string
	DbName string

	db *sql.DB
	initialized bool
}

func (c *Config) Init() error {
	return c.InitWithDb(c.DbName)
}

func (c *Config) InitWithDb(name string) error {
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s sslmode=disable`,
		c.Host, c.Port, c.User)
	if name != `` {
		psqlInfo = psqlInfo + fmt.Sprintf(` dbname=%s`, name)
	}
	log.Info().Str(`connection`, psqlInfo).Msg(`connecting to database`)
	db, err := sql.Open(`postgres`, psqlInfo)
	if err != nil {
		return fmt.Errorf(`failed to connect to database: %w`, err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf(`failed to ping database: %w`, err)
	}
	c.db = db
	c.initialized = true
	log.Info().Msg(`connected to database`)
	return nil
}

func (c *Config) Close() error {
	return c.db.Close()
}

func (c *Config) UseDatabase(ctx context.Context, name string) error {
	if !c.initialized {
		return ErrNotInitialized
	}
	return c.ExecContext(ctx, fmt.Sprint(`USE DATABASE `), name)
}

func (c *Config) CreateDatabase(ctx context.Context, name string) error {
	return c.ExecContext(ctx, fmt.Sprint(`CREATE DATABASE IF NOT EXISTS `, name))
}

func (c *Config) CreateTable(ctx context.Context, name string) error {
	return c.ExecContext(ctx, fmt.Sprint(`CREATE TABLE IF NOT EXISTS `, name))
}

func (c *Config) CreateDBUser(ctx context.Context, name string) error {
	return c.ExecContext(ctx, `CREATE USER IF NOT EXISTS $1`, name)
}

func (c *Config) DropDatabase(ctx context.Context, name string) error {
	return c.ExecContext(ctx, fmt.Sprint(`DROP DATABASE IF EXISTS `, name))
}

func (c *Config) ExecContext(ctx context.Context, stmt string, args ...interface{}) error {
	_, err := c.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		log.Error().Err(err).Str(`stmt`, stmt).Msg(`failed to execute statement`)
	}
	return err
}

func (c *Config) QueryContext(ctx context.Context, stmt string, args ...interface{}) (*sql.Rows, error) {
	rows, err := c.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		log.Error().Err(err).Str(`stmt`, stmt).Msg(`failed to execute query`)
		return nil, err
	}
	return rows, nil
}