package todo

import (
	"context"
	"database/sql"

	"brice.io/todo/core/entities"
)

// ToDo Repository
type ToDoStorage interface {
	insertToDoInDb(td *entities.ToDo)
	listToDoInDb() []entities.ToDo
}

// ToDo Service
type ToDoService struct {
	db *sql.DB
	ctx context.Context
}

// Get ToDo list
func (t *ToDoService) listToDoInDb() []entities.ToDo {
	return nil
}

// Insert New ToDo
func (t *ToDoService) insertToDoInDb(td *entities.ToDo) {

}

// Constructor
func NewToDoStorage(ctx context.Context) ToDoStorage {
	return &ToDoService{ctx: ctx}
}