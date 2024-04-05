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
	DB_ENGINE   string
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USERNAME string
}

func GetEnv(env_file string) AppEnv {
	err := godotenv.Load(env_file)
	elog.New(elog.PANIC, "Error loading "+env_file+"file", err)

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	return AppEnv{
		SERVER_PORT: port,
		DB_ENGINE: os.Getenv("DB_ENGINE"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
	}
}
