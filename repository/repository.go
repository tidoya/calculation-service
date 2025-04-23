package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
