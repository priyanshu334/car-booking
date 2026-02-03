package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/priyanshu334/go_car_book/internal/modules/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(name, email, password string) error
	Login(email, password string) (string, string, error)
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

func (s *service) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credintials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		return "", "", errors.New("invalid credintials")
	}
	access, err := auth.GenerateAccessToken(user.ID.String(), string(user.Role))
	if err != nil {
		return "", "", err
	}

	refresh, err := auth.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return "", "", err
	}

	return access, refresh, err

}
