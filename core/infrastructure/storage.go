package todo

import (
	"context"
	"database/sql"

	"brice.io/todo/core/entities"
	"brice.io/todo/internal/helpers/elog"
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
	querystring := `SELECT id, name, complete FROM todo ORDER BY _name ASC`
	rows, err := t.db.QueryContext(t.ctx, querystring)

	// iferr
	if err != nil {
		return nil
	}

	var todos []entities.ToDo
	for rows.Next() {
		var td entities.ToDo
		err = rows.Scan(&td.ID, &td.Name, &td.Complete)
		go elog.New(elog.ERROR, "Error getting list of ToDo", err)

		todos = append(todos, td)
	}
	defer rows.Close()

	return todos
}

// Insert New ToDo
func (t *ToDoService) insertToDoInDb(td *entities.ToDo) {

}

// Constructor
func NewToDoStorage(ctx context.Context) ToDoStorage {
	return &ToDoService{ctx: ctx}
}