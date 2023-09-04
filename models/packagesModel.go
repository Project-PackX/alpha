package models

import "gorm.io/gorm"

// Package DB table model
type Package struct {
	gorm.Model
	Id                  uint    // Entity id
	UserID              uint    // Sender's id
	ReceiverId          uint    // Receiver's id
	DestinationLockerId uint    // Destination locker's id
	Size                string  // Small, Medium, Large, needs to be an enum
	Price               float64 // Delivery fee
	Code                string  // Code to open the locker
	Note                string  // Extra note
	CourierID           uint    // Courier's id who delivers the package
}

// Name of the Package structs in the DB
func (Package) TableName() string {
	return "public.packages"
}
