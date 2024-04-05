package todo

import "brice.io/todo/core/entities"

// Gateway access to storage
type ToDoGateway interface {
	CreateToDo(td *entities.ToDo) (int, entities.Response)
	ListToDo() (int, []entities.ToDo)
}

// Domain Logic
type ToDoLogic struct {
	St ToDoStorage
}

