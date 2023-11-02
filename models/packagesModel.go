package models

import (
	"time"

	"gorm.io/gorm"
)

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
	TrackID             string    // Code for package tracking
	UserID              uint      // Sender user's id
	SenderLockerId      uint      // Sender locker's id
	DestinationLockerId uint      // Destination locker's id
	ReceiverName        string    // Name of the person who gets the package
	ReceiverEmail       string    // Email of the person who gets the package
	Size                string    // Could be small, medium or large
	DeliverySpeed       string    // Type of delivery
	Price               float64   // Delivery fee (how to calculate?)
	Code                string    // Code to open the locker (both sender and receiver) - random 6 digit number maybe?
	DeliveryDate        time.Time // Date when the package arrives
	Note                string    // Extra note
	CourierID           uint      // Courier's id who delivers the package (can be different from send. locker to warehouse and warehouse to dest. locker?)
}

// Name of the Package structs in the DB
func (Package) TableName() string {
	return "public.packages"
}
