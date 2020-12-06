package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PQHandler struct {
	*SQLHandler
}

func NewPQHandler(connect string) (*PQHandler, error) {
	db, err := sql.Open("postgres", connect)
	return &PQHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
