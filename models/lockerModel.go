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
	/* Possibly there will be 2 more column for X and Y geo. points if we want to show them on a map */

	Capacity uint // Shows how many packages could be there
}

// Name of the Locker structs in the DB
func (Locker) TableName() string {
	return "public.lockers"
}
