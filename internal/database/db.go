package database

import (
	"database/sql"
	"time"

	"brice.io/todo/internal/helpers/elog"
)

type Database interface {
	ConnectDB() *sql.DB
}

// Max Seconds before attempting to reconnect
const DB_CONNECTION_TIMEOUT = 10

// Attempt db connection
func try(err error, db *sql.DB, counts *int) error {
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
