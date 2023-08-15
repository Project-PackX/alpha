package models

import "gorm.io/gorm"

// Programon belül a Csomag típus mezői
type Courier struct {
	gorm.Model
	Name  string
	Phone string
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (Courier) TableName() string {
	return "public.couriers"
}
