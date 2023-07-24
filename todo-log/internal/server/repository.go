package server

import "github.com/jmoiron/sqlx"

type Todo struct {
	db *sqlx.DB
}
