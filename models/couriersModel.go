package models

import "gorm.io/gorm"

// Courier DB table model
type Courier struct {
	gorm.Model
	/*
			Inside gorm.Model:

			ID        uint           `gorm:"primaryKey"` incrementing
		    CreatedAt time.Time
		    UpdatedAt time.Time
		    DeletedAt gorm.DeletedAt `gorm:"index"`
	*/
	Name  string // Name of the courier
	Phone string // Phone address of the courier
}

// Name of the Courier structs in the DB
func (Courier) TableName() string {
	return "public.couriers"
}
