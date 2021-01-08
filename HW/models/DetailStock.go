package models

type DetailStock struct {
	DetailID uint `gorm:"detail_id"`
	StockID  uint `gorm:"supplier_id"`
	Count    uint `gorm:"count"`
}
