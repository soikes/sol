package crdb

import (
	"database/sql"
)

const (
	host = `localhost`
	port = `26257`
	user = `sol`
	password = `sol`
	db = `sol`
)

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	db *sql.DB
}

func (c *Config) Init() error {
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s sslmode=disable`,
		host, port, user, password, db)
	db, err := sql.Open(`postgres`, psqlInfo)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	c.db = db
}

func (c *Config) Shutdown() error {
	return c.db.Close()
}