package models

import "gorm.io/gorm"

type DetailStock struct {
	gorm.Model
	DetailID uint `gorm:"detail_id"`
	StockID  uint `gorm:"supplier_id"`
	Count    uint `gorm:"count"`
}
