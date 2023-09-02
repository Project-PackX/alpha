package models

// LockerGroup Table Model
type LockerGroup struct {
	ID      uint     // Identification of the locker group
	City    string   // For human readability, which city is this group in
	Lockers []Locker // List of the lockers, which are belongs to this group
}

// Giving the scema and table name for the 'LockerGroup' model
func (LockerGroup) TableName() string {
	return "public.lockergroups"
}
