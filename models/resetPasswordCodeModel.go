package models

import (
	"gorm.io/gorm"
)

// ResetPasswordCode DB table model
type ResetPasswordCode struct {
	gorm.Model
	/*
			Inside gorm.Model:

			ID        uint           `gorm:"primaryKey"` incrementing
		    CreatedAt time.Time
		    UpdatedAt time.Time
		    DeletedAt gorm.DeletedAt `gorm:"index"`
	*/
	Code    string
	User_id uint
}

// Name of the ResetPasswordCode structs in the DB
func (ResetPasswordCode) TableName() string {
	return "public.reset_password_code"
}
