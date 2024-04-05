package cmd

import (
	"context"

	"brice.io/todo/env"
	"brice.io/todo/internal/database"
	"brice.io/todo/server"
)

func Start() {
	
	// ctx
	ctx := context.Background()

	// env
	_env := env.GetEnv(".env.development")

	// DB
	db := database.NewMySQLDatabase(ctx, _env).ConnectDB()
	defer db.Close()

	// Server
	server.NewEchoServer(ctx, _env.SERVER_PORT).Run()

}
