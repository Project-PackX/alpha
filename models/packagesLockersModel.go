package models

// PackageLocker DB table model
type PackageLocker struct {
	Package_id uint `gorm:"primaryKey"`
	Locker_id  uint `gorm:"primaryKey"`
}

// Name of the PackageStatus structs in the DB
func (PackageLocker) TableName() string {
	return "public.packageslockers"
}
