package repository

import (
	Calculation "calculation-service/internal/entity/user"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user Calculation.User) (int, error)
	GetUser(username, password string) (Calculation.User, error)
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
