package car

import "github.com/google/uuid"

type Service interface {
	AddCar(ownerID uuid.UUID, name, brand string, price int) error
	ListCars() ([]Car, error)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

func (s *service) AddCar(ownerID uuid.UUID, name, brand string, price int) error {
	car := &Car{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Name:        name,
		Brand:       brand,
		PricePerDay: price,
		Status:      CarPending,
	}
	return s.repo.Create(car)
}

func (s *service) ListCars() ([]Car, error) {
	return s.repo.FindAllActive()
}
