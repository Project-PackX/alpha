package models

import (
	"gorm.io/gorm"
)

// Programon belül a Csomag típus mezői
type User struct {
	gorm.Model
	/*
			Benne van:

			ID        uint           `gorm:"primaryKey"`
		    CreatedAt time.Time
		    UpdatedAt time.Time
		    DeletedAt gorm.DeletedAt `gorm:"index"`
	*/

	Name     string
	Address  string
	Phone    string
	Email    string
	Packages []Package
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (User) TableName() string {
	return "public.users"
}
