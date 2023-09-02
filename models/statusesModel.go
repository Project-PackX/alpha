package models

// Status DB table model
type Status struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

// Name of the Status structs in the DB
func (Status) TableName() string {
	return "public.statuses"
}
