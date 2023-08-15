package models

import "gorm.io/gorm"

// Programon belül a Csomag típus mezői
type Package struct {
	gorm.Model
	UserID             uint    // Küldő fél (feladó) azonosítója
	DestinationAddress string  // A kiszállítási cím
	Content            string  // A csomag tartalma
	Price              float64 // A csomag küldésének ára
	Note               string  // Egyéb megjegyzés, opcionális
	CourierID          uint    // A csomagot szállító futár azonosítója
}

// Megadjuk a séma és táblanevét a "Package"-eket tartalmazó adattáblának
func (Package) TableName() string {
	return "public.packages"
}
