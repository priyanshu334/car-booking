package booking

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	BookingPending   BookingStatus = "PENDING"
	BookingConfirmed BookingStatus = "CONFIRMED"
	BookingCancelled BookingStatus = "CANCELLED"
	BookingCompleted BookingStatus = "COMPLETED"
)

type Booking struct {
	ID        uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID     `gorm:"type:uuid;not null"`
	CarID     uuid.UUID     `gorm:"type:uuid;not null"`
	StartDate time.Time     `gorm:"not null"`
	EndDate   time.Time     `gorm:"not null"`
	Status    BookingStatus `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
