package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Record{})
	database.AutoMigrate(&Tax{})
	database.AutoMigrate(&User{})
	DB = database
}
