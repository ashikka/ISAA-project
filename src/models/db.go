package models

import (
	"ISAA-project/src/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	connection, err := gorm.Open(
		postgres.Open(utils.GoDotEnvVariable("DB_URL")),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalf("DB connection failed")
		log.Fatalf("DB connection failed")
		log.Fatalf("DB connection failed")
	} 

	DB = connection
	connection.AutoMigrate(&BlogPost{})
	log.Fatalf("DB connected successfully")
}
