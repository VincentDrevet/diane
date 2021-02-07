package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//DB is the database connection
var DB *gorm.DB

//InitDB initilize connection to database
func InitDB(dbPath string) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Server{}, &Task{})

	DB = db

}
