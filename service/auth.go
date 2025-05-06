package service

import (
	Calculation "calculation-service/internal/entity/user"
	"calculation-service/repository"
	"crypto/sha1"
	"fmt"
)

const (
	salt = "asjdhaskdhqwhejkqwhe123123hhhrjqwehqjw"
)

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user Calculation.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
