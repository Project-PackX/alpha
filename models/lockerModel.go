package models

import "gorm.io/gorm"

// Locker DB table Model
type Locker struct {
	gorm.Model
	Address       string // Address of the locker
	LockerGroupID uint   // A Lockergroup ID where the locker belongs to
}

// Name of the Locker structs in the DB
func (Locker) TableName() string {
	return "public.lockers"
}
