package models

import "gorm.io/gorm"

// Package DB table model
type Package struct {
	gorm.Model
	UserID             uint    // Sender's id
	DestinationAddress string  // Destination address (street + house number)
	Content            string  // What is inside the package
	Price              float64 // Delivery fee
	Note               string  // Extra note
	CourierID          uint    // Courier's id who delivers the package
}

// Name of the Package structs in the DB
func (Package) TableName() string {
	return "public.packages"
}
