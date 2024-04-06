package server

import app "brice.io/todo/core/application"

// ToDo routes
func (es *EchoServer) toDoRoutes() {
	// call ToDo HTTP Service
	todo := app.NewToDoHTTPService(es.ctx, es.db)

	es.GET("/todo", todo.ListHandler)
	es.POST("/todo/create", todo.CreateHandler)
}

// Other routes
// TODO: user management, auth

// All routes
func (es *EchoServer) routes() {
	es.toDoRoutes()
}