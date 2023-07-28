package server

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	api "github.com/lathief/simple-grpc-go/todo-log/api/v1"
	"log"
)

type TodoRepo struct {
	db *sqlx.DB
}

type TodoRepoInterface interface {
	Get(id int) (*api.Todo, error)
	GetAll() (*[]api.Todo, error)
	Insert(todo *api.Todo) (*api.Todo, error)
	Update(todo *api.Todo) (*api.Todo, error)
	Delete(id int) error
}

func NewTodo(db *sqlx.DB) (*TodoRepo, error) {
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
	return &TodoRepo{
		db: db,
	}, nil
}

func (t *TodoRepo) Get(id int) (*api.Todo, error) {
	log.Printf("Getting %d", id)
	var todo = api.Todo{}
	row := t.db.QueryRow("SELECT * FROM todos WHERE id=?", id)
	if err := row.Scan(&todo.Id, &todo.CreatedAt, &todo.Title, &todo.Description, &todo.Status); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return &api.Todo{}, err
	}
	return &todo, nil
}
func (t *TodoRepo) GetAll() ([]*api.Todo, error) {
	log.Printf("Get all")
	rows, err := t.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*api.Todo
	for rows.Next() {
		todo := api.Todo{}
		err = rows.Scan(&todo.Id, &todo.CreatedAt, &todo.Title, &todo.Description, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, err
}
func (t *TodoRepo) Insert(todo *api.Todo) (*api.Todo, error) {
	log.Printf("Insert")
	res, err := t.db.Exec("INSERT INTO todos(title, description, status) VALUES(?,?,?);", todo.Title, todo.Description, todo.Status)
	if err != nil {
		return &api.Todo{}, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return &api.Todo{}, err
	}
	todo.Id = int32(id)
	log.Printf("Added %v as %d", todo, id)
	return todo, nil
}
func (t *TodoRepo) Update(todo *api.Todo) (*api.Todo, error) {
	log.Printf("Update")
	res, err := t.db.Exec(
		`UPDATE todos SET title=?, description=?, 
				status=? WHERE id = :id`,
		todo.Title, todo.Description, todo.Status)
	if err != nil {
		return &api.Todo{}, err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &api.Todo{}, sql.ErrNoRows
	}
	row := t.db.QueryRow("SELECT * FROM todos WHERE id=?", todo.Id)
	if err = row.Scan(&todo.Id, &todo.CreatedAt, &todo.Title, &todo.Description, &todo.Status); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return &api.Todo{}, err
	}
	return todo, nil
}
func (t *TodoRepo) Delete(id int) error {
	res, err := t.db.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	log.Printf("Delete %d", id)
	return nil
}
