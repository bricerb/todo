package app

import (
	"context"

	"brice.io/todo/core/entities"
	todo "brice.io/todo/core/infrastructure"
	"brice.io/todo/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// HTTP Handler for ToDo
type ToDoHTTPService struct {
	gtw todo.ToDoGateway
}

func (t *ToDoHTTPService) ListHandler(c echo.Context) error {
	status, res := t.gtw.ListToDo()
	return c.JSON(status, res)
}

func (t *ToDoHTTPService) CreateHandler(c echo.Context) error {

	// New entity
	td := new(entities.ToDo)
	// Bind data to struct
	c.Bind(td)
	// validation

	// Generate new UUID
	td.ID = helpers.UUID()

	// Default value
	td.Complete = false

	status, res := t.gtw.CreateToDo(td)
	return c.JSON(status, res)
}

// Constructor
func NewToDoHTTPService(ctx context.Context, db *sqlx.DB) *ToDoHTTPService {
	return &ToDoHTTPService{todo.NewToDoGateway(ctx, db)}
}