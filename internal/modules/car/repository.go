package car

import "gorm.io/gorm"

type Repository interface {
	Create(car *Car) error
	FindAllActive() ([]Car, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(car *Car) error {
	return r.db.Create(car).Error

}

func (r *repository) FindAllActive() ([]Car, error) {
	var cars []Car
	err := r.db.Where("status=?", CarActive).Find(&cars).Error
	return cars, err
}
