package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Address string `gorm:"address"`
	Phone   string `gorm:"phone"`
}

type StockView struct {
	Stock
	Count uint `gorm:"count"`
}
