package database

import (
	"time"

	"brice.io/todo/internal/helpers/elog"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	ConnectDB() *sqlx.DB
}

// Max Seconds before attempting to reconnect
const DB_CONNECTION_TIMEOUT = 10

// Attempt db connection
func try(err error, db *sqlx.DB, counts *int) error {
	if err != nil {
		// increasing counter
		elog.New(elog.ERROR, "Trying to connect to database", err)
		*counts++

		// Connection Fails
		if *counts > DB_CONNECTION_TIMEOUT {
			elog.New(elog.PANIC, "Failure to connect to database", err)
		}

		// Log and retry
		elog.New(elog.ERROR, "Retrying", err)
		time.Sleep(time.Second)
	}

	return err
}
