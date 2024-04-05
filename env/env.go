package env

import (
	"os"

	"brice.io/todo/internal/helpers/elog"
	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "80"

type AppEnv struct {
	// Server Envs
	SERVER_PORT string

	// DB Envs

}

func GetEnv(env_file string) AppEnv {
	err := godotenv.Load(env_file)
	elog.New(elog.PANIC, "Error loading " + env_file + "file", err)

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	return AppEnv{
		SERVER_PORT: port,
	}
}