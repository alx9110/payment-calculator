package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(connection_string interface{}) {
	database, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s", connection_string)))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Record{})
	database.AutoMigrate(&Tax{})
	database.AutoMigrate(&User{})
	DB = database
}
