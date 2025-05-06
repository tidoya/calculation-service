package service

import (
	Calculation "calculation-service/internal/entity/user"
	"calculation-service/repository"
)

type Autorization interface {
	CreateUser(user Calculation.User) (int, error)
}

type Service struct {
	Autorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
