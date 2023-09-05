package models

// PackageStatus DB table model
type PackageStatus struct {
	Package_id uint `gorm:"primaryKey"`
	Status_id  uint `gorm:"primaryKey"`
}

// Name of the PackageStatus structs in the DB
func (PackageStatus) TableName() string {
	return "public.packagestatuses"
}
