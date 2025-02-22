package config

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresConfig struct {
	Server   string
	Username string
	Password string
	Database string
}

var (
	db *bun.DB
)

func NewConnection(c *PostgresConfig) (e error) {

	ds := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Server, c.Database)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(ds)))

	e = sqldb.Ping()
	if e == nil {
		db = bun.NewDB(sqldb, pgdialect.New())

		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))

		slog.Info(fmt.Sprintf("Connected to Postgres Server: %s@%s", c.Server, c.Database))
	}

	return
}

func GetDB() *bun.DB {
	return db
}
