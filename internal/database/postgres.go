package database

import (
	"context"
	"database/sql"
	"fmt"

	"brice.io/todo/env"
	_ "github.com/lib/pq"
)

// DB Class
type PostgresDB struct {
	ctx    context.Context
	config env.AppEnv
}

// Ping
func (m *PostgresDB) ping(err error, db *sql.DB) error {
	if err != nil {
		return err
	}

	// try ping
	err = db.PingContext(m.ctx)
	if err != nil {
		return err
	}

	return nil
}

// Connect
func (m *PostgresDB) connect() (*sql.DB, error) {
	// try open connection
	db, err := sql.Open(
		m.config.DB_ENGINE,
		fmt.Sprintf(
			"%s@tcp(%s:%s)/%s",
			m.config.DB_USERNAME,
			m.config.DB_HOST,
			m.config.DB_PORT,
			m.config.DB_DATABASE,
		))

	// try ping
	if m.ping(err, db) != nil {
		return nil, err
	}

	return db, nil
}

// Get connection
func (m *PostgresDB) ConnectDB() *sql.DB {
	counts := 0

	for {
		db, err := m.connect()

		if try(err, db, &counts) == nil {
			return db
		}
		continue
	}
}

// Constructor
func NewMySQLDatabase(ctx context.Context, ec env.AppEnv) Database {
	return &PostgresDB{ctx: ctx, config: ec}
}