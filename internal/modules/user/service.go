package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(name, email, password string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(name, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     RoleCustomer,
	}

	if err := s.repo.Create(user); err != nil {
		return errors.New("email already exists")
	}

	return nil
}
