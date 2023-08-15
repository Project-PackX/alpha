package models

// Programon belül a Csomag típus mezői
type PackageStatus struct {
	Package_id uint `gorm:"primaryKey"`
	Status_id  uint `gorm:"primaryKey"`
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (PackageStatus) TableName() string {
	return "public.packagestatuses"
}
