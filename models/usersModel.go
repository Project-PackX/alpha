package models

import (
	"gorm.io/gorm"
)

// User DB table model
type User struct {
	gorm.Model
	/*
			Inside gorm.Model:

			ID        uint           `gorm:"primaryKey"` incrementing
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

// Name of the User structs in the DB
func (User) TableName() string {
	return "public.users"
}
