package models

// Programon belül a Csomag típus mezői
type Status struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (Status) TableName() string {
	return "public.statuses"
}
