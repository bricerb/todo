package todo

import (
	"context"
	"fmt"

	"brice.io/todo/core/entities"
	"brice.io/todo/internal/helpers/elog"
	"github.com/jmoiron/sqlx"
)

// ToDo Repository
type ToDoStorage interface {
	insertToDoInDb(td *entities.ToDo)
	listToDoInDb() []entities.ToDo
}

// ToDo Service
type ToDoService struct {
	db  *sqlx.DB
	ctx context.Context
}

// Get ToDo list
func (t *ToDoService) listToDoInDb() []entities.ToDo {
	querystring := `SELECT id, name, complete FROM todo ORDER BY name ASC`
	rows, err := t.db.QueryxContext(t.ctx, querystring)

	// iferr
	if err != nil {
		elog.New(elog.ERROR, "Error reading ToDos from db", err)
		return nil
	}

	var todos []entities.ToDo
	for rows.Next() {
		var td entities.ToDo
		err = rows.StructScan(&td)
		if err != nil {
			go elog.New(elog.ERROR, "Error getting list of ToDo", err)
		}
		

		todos = append(todos, td)
	}
	defer rows.Close()

	return todos
}

// Insert New ToDo
func (t *ToDoService) insertToDoInDb(td *entities.ToDo) {
	querystring := `INSERT INTO todo (id, name, complete) VALUES ($1, $2, $3)`
	_, err := t.db.ExecContext(t.ctx, querystring, td.ID, td.Name, td.Complete)
	go elog.New(elog.ERROR, "Error inserting a ToDo in db", err)
}

// Constructor
func NewToDoStorage(ctx context.Context, db *sqlx.DB) ToDoStorage {
	return &ToDoService{db: db, ctx: ctx}
}
