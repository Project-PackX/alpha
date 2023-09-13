package models

import (
	"time"

	"gorm.io/gorm"
)

// Locker DB table Model
type Locker struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Address   string         // Address of the locker
}

// Name of the Locker structs in the DB
func (Locker) TableName() string {
	return "public.lockers"
}
