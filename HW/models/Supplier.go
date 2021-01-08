package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string `gorm:"name"`
	Address string `gorm:"address"`
	Phone   string `gorm:"phone"`
}
