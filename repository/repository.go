package repository

import (
	Calculation "calculation-service/internal/entity/user"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user Calculation.User) (int, error)
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
