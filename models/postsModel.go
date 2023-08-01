package models

import (
	"gorm.io/gorm"
)

// Programon belül a Csomag típus mezői
type Package struct {
	gorm.Model // ID és időbélyegek benne vannak ebben
	Sender     string
	Price      float32
	Delivered  bool
}
