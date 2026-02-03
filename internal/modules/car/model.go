package car

import (
	"time"

	"github.com/google/uuid"
)

type CarStatus string

const (
	CarPending  CarStatus = "PENDING"
	CarActive   CarStatus = "ACTIVE"
	CarInactive CarStatus = "INACTIVE"
)

type Car struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"not null"`
	Brand       string    `gorm:"not null"`
	PricePerDay int       `gorm:"not null"`
	Status      CarStatus `gorm:"type:varchar(20);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

