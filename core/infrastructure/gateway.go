package todo

import (
	"context"
	"database/sql"
	"net/http"

	"brice.io/todo/core/entities"
)

// Gateway access to storage
type ToDoGateway interface {
	CreateToDo(td *entities.ToDo) (int, entities.Response)
	ListToDo() (int, []entities.ToDo)
}

// Domain Logic
type ToDoLogic struct {
	St ToDoStorage
}

// List ToDo
func (t *ToDoLogic) ListToDo() (int, []entities.ToDo) {
	return http.StatusOK, t.St.listToDoInDb()
}

// Create ToDo
func (t *ToDoLogic) CreateToDo(td *entities.ToDo) (int, entities.Response) {
	
	// Invalid name logic
	if td.Name == "" {
		return http.StatusOK, entities.Response{
			Message: "Invalid ToDo",
			Success: false,
		}
	}

	// Create ToDo
	go t.St.insertToDoInDb(td)
	
	// Accepted response
	return http.StatusAccepted, entities.Response{
		Message: "ToDo successfully added",
		Success: true,
	}
}

// Constructor
func NewToDoGateway(ctx context.Context, db *sql.DB) ToDoGateway {
	return &ToDoLogic{NewToDoStorage(ctx, db)}
}