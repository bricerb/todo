package cmd

import (
	"context"
	"log"

	"brice.io/todo/env"
	"brice.io/todo/internal/database"
	"brice.io/todo/server"
)

func Start() {
	
	// ctx
	ctx := context.Background()

	// env
	_env := env.GetEnv(".env.development")

	// log env
	log.Printf("Configs: %v", _env)

	// DB
	db := database.NewMySQLDatabase(ctx, _env).ConnectDB()
	defer db.Close()

	// Log start info
	log.Printf("db: %v", db)

	// Server
	server.NewEchoServer(ctx, db, _env.SERVER_PORT).Run()

}
