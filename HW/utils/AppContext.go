package utils

import "gorm.io/gorm"

type AppContext struct {
	Config Config
	DB     *gorm.DB
}
