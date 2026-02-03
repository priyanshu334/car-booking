package booking

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Service interface {
	CreateBooking(userID, cardID uuid.UUID, start, end time.Time) error
}

type service struct {
	repo Repository
	db   *gorm.DB
}

func NewService(repo Repository, db *gorm.DB) *service {
	return &service{repo: repo, db: db}
}

func (s *service) CreateBooking(userID, carID uuid.UUID, start, end time.Time) error {
	if end.Before(start) {
		return errors.New("invalid date range")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		overlap, err := s.repo.HasOverlap(carID.String(), start, end)
		if err != nil {
			return err
		}

		if overlap {
			return errors.New("Car already booked for selected dates")

		}
		booking := &Booking{
			ID:        uuid.New(),
			UserID:    userID,
			CarID:     carID,
			StartDate: start,
			EndDate:   end,
			Status:    BookingPending,
		}

		return s.repo.Create(tx, booking)
	})
}
