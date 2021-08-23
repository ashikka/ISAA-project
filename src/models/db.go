package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func ConnectDB(dbURL string) *gorm.DB {

	db, err = gorm.Open(dbURL)

	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	defer db.Close()

	return db.AutoMigrate(&BlogPost{})
}
