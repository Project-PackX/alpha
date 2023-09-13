package models

// LockerGroup DB table Model
type LockerGroup struct {
	ID   uint   // Identification of the locker group
	City string // For human readability, which city is this group in
}

// Name of the LockerGroup structs in the DB
func (LockerGroup) TableName() string {
	return "public.lockergroups"
}
