package models

import (
	"gorm.io/gorm"
)

// Locker DB table Model
type Locker struct {
	gorm.Model
	/*
		ID        string `gorm:"primaryKey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	*/
	City      string //Name of the city where the locker is
	Address   string // Address of the locker
	Latitude  float64
	Longitude float64

	Capacity uint // Shows how many packages could be there
}

// Name of the Locker structs in the DB
func (Locker) TableName() string {
	return "public.lockers"
}
