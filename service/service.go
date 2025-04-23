package service

import "calculation-service/repository"

type Autorization interface {
}

type Service struct {
	Autorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
