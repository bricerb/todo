package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	*echo.Echo
	ctx context.Context
	port string
}

func (es *EchoServer) configure() {

	// console output
	es.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_custom} : ${method} => ${uri}, status={$status} ::$error}\n",
		CustomTimeFormat: "15:04.05.00000",
	}))

	// graceful recover from panic
	es.Use(middleware.Recover())

	// cors
	es.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// auth

}

// Run Echo Server
func (es *EchoServer) Run() {
	es.Logger.Fatal(es.Start(":" + es.port))
}

// New Server instance
func NewEchoServer(ctx context.Context, app_port string) Server {
	if app_port == "" {
		app_port = "8080"
	}

	server := &EchoServer{
		echo.New(),
		ctx,
		app_port,
	}
	server.configure()
	server.routes()

	return server
}