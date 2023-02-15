package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func SetDatabase() {

	dsn := "host=localhost user=postgres password=fred port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error setting DB")
	}

	err = db.AutoMigrate(&Location{})
	if err != nil {
		return
	}

	Database = db
}
