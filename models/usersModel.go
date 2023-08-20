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
	Password string
	Packages []Package
}

// UserRequest DTO for incoming request bodies
type UserRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (User) TableName() string {
	return "public.users"
}
