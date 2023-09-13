package models

import "gorm.io/gorm"

// Package DB table model
type Package struct {
	gorm.Model
	/*
			Inside gorm.Model:

			ID        uint           `gorm:"primaryKey"` incrementing
		    CreatedAt time.Time
		    UpdatedAt time.Time
		    DeletedAt gorm.DeletedAt `gorm:"index"`
	*/
	UserID              uint    // Sender user's id
	SenderLockerId      string  // Sender locker's id
	DestinationLockerId string  // Destination locker's id
	Size                string  // Could be small, medium or large
	Price               float64 // Delivery fee (how to calculate?)
	Code                uint    // Code to open the locker (both sender and receiver) - random 6 digit number maybe?
	Note                string  // Extra note
	CourierID           uint    // Courier's id who delivers the package (can be different from send. locker to warehouse and warehouse to dest. locker?)
}

// Name of the Package structs in the DB
func (Package) TableName() string {
	return "public.packages"
}
