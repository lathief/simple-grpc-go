package server

import (
	"github.com/jmoiron/sqlx"
	api "github.com/lathief/simple-grpc-go/todo-log/api/v1"
	"log"
)

type Todo struct {
	db *sqlx.DB
}

type TodoInterface interface {
	Get(id int) (*api.Todo, error)
	GetAll() (*[]api.Todo, error)
	Insert(todo *api.Todo) (*api.Todo, error)
	Update(todo *api.Todo) (*api.Todo, error)
	Delete(id int) error
}

func NewTodo(db *sqlx.DB) (*Todo, error) {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY,
		created_at TIMESTAMP NOT NULL,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		status VARCHAR(50) NOT NULL
		);`,
	); err != nil {
		return nil, err
	}
	return &Todo{
		db: db,
	}, nil
}

func (t *Todo) Get(id int) (*api.Todo, error) {
	log.Printf("Getting %d", id)
	var todo = api.Todo{}
	err := t.db.Get(todo, "SELECT * FROM todos WHERE id=?", id)
	return &todo, err
}
func (t *Todo) GetAll() (*[]api.Todo, error) {
	log.Printf("Get all")
	var todos []api.Todo
	err := t.db.Select(todos, "SELECT * FROM todos")
	return &todos, err
}
