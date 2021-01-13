package utils

import (
	"../models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDatabase(config Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatalln("ERROR(failed to connect database): " + err.Error())
	}

	err = db.AutoMigrate(
		&models.Detail{},
		&models.Supplier{},
		&models.DetailStock{},
		&models.Stock{},
	)
	if err != nil {
		log.Fatalln("ERROR(failed to migrate database): " + err.Error())
	}
	return db
}
