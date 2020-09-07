package crdb

import (
	"context"
	"errors"
	"fmt"

	"soikke.li/sol"
	"soikke.li/sol/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ErrNotInitialized = errors.New(`db has not been initialized`)
	ErrNoRows         = errors.New(`no rows in result set`)
)

type Config struct {
	sol.Component

	Host string
	Port string
	User string
	DbName string

	db *sqlx.DB
	initialized bool
}

func (cfg *Config) Init(log log.Logger) {
	cfg.Component.Init(`crdb`, log)
}

func (cfg *Config) InitDB() error {
	return cfg.InitWithDb(cfg.DbName)
}

func (cfg *Config) InitWithDb(name string) error {
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s sslmode=disable`,
		cfg.Host, cfg.Port, cfg.User)
	if name != `` {
		psqlInfo = psqlInfo + fmt.Sprintf(` dbname=%s`, name)
	}
	cfg.Log.Info().Str(`connection`, psqlInfo).Msg(`connecting to database`)
	db, err := sqlx.Open(`postgres`, psqlInfo)
	if err != nil {
		return fmt.Errorf(`failed to connect to database: %w`, err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf(`failed to ping database: %w`, err)
	}
	cfg.db = db
	cfg.initialized = true
	cfg.Log.Info().Msg(`connected to database`)
	return nil
}

func (cfg *Config) Close() error {
	return cfg.db.Close()
}

func (cfg *Config) UseDatabase(ctx context.Context, name string) error {
	if !cfg.initialized {
		return ErrNotInitialized
	}
	return cfg.ExecContext(ctx, fmt.Sprint(`USE DATABASE `), name)
}

func (cfg *Config) CreateDatabase(ctx context.Context, name string) error {
	return cfg.ExecContext(ctx, fmt.Sprint(`CREATE DATABASE IF NOT EXISTS `, name))
}

func (cfg *Config) CreateTable(ctx context.Context, name string) error {
	return cfg.ExecContext(ctx, fmt.Sprint(`CREATE TABLE IF NOT EXISTS `, name))
}

func (cfg *Config) CreateDBUser(ctx context.Context, name string) error {
	return cfg.ExecContext(ctx, `CREATE USER IF NOT EXISTS $1`, name)
}

func (cfg *Config) DropDatabase(ctx context.Context, name string) error {
	return cfg.ExecContext(ctx, fmt.Sprint(`DROP DATABASE IF EXISTS `, name))
}

func (cfg *Config) ExecContext(ctx context.Context, stmt string, args ...interface{}) error {
	_, err := cfg.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		cfg.Log.Error().Err(err).Str(`stmt`, stmt).Msg(`failed to execute statement`)
	}
	return err
}

func (cfg *Config) QueryContext(ctx context.Context, stmt string, args ...interface{}) (*sqlx.Rows, error) {
	rows, err := cfg.db.QueryxContext(ctx, stmt, args...)
	if err != nil {
		cfg.Log.Error().Err(err).Str(`stmt`, stmt).Msg(`failed to execute query`)
		return nil, err
	}
	return rows, nil
}