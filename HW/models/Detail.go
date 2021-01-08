package models

import "gorm.io/gorm"

type Detail struct {
	gorm.Model
	Name        string  `gorm:"name"`
	Description string  `gorm:"description"`
	Weight      float64 `gorm:"weight"`
	Image       string  `gorm:"image"`
	Supplier    uint    `gorm:"supplier_id"`
}
