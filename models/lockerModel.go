package models

import "gorm.io/gorm"

// Locker Table Model
type Locker struct {
	gorm.Model
	Address       string // Address of the locker
	LockerGroupID uint   // A Lockergroup ID where the locker belongs to
}

// Giving the scema and table name for the 'Locker' model
func (Locker) TableName() string {
	return "public.lockers"
}
