package database

import (
	"context"
	"fmt"

	"brice.io/todo/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB Class
type PostgresDB struct {
	ctx    context.Context
	config env.AppEnv
}

// Ping
func (m *PostgresDB) ping(err error, db *sqlx.DB) error {
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
func (m *PostgresDB) connect() (*sqlx.DB, error) {
	// try open connection
	conn_str := fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s", m.config.DB_USERNAME, m.config.DB_DATABASE, m.config.DB_HOST)
	db, err := sqlx.Connect(m.config.DB_ENGINE, conn_str)

	// try ping
	if m.ping(err, db) != nil {
		return nil, err
	}

	return db, nil
}

// Get connection
func (m *PostgresDB) ConnectDB() *sqlx.DB {
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
