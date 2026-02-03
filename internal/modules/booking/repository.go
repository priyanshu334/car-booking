package booking

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	HasOverlap(carID string, start, end time.Time) (bool, error)
	Create(tx *gorm.DB, booking *Booking) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) HasOverlap(carID string, start, end time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&Booking{}).
		Where("car_id = ?", carID).
		Where("status IN ?", []BookingStatus{BookingPending, BookingConfirmed}).
		Where("start_date <= ? AND end_date >= ?", end, start).
		Count(&count).Error

	return count > 0, err
}

func (r *repository) Create(tx *gorm.DB, booking *Booking) error {
	return tx.Create(booking).Error
}
