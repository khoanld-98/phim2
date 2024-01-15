package main

import (
	"web/models"
	"web/route"
	"web/store"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load(".env")
	db = store.ConnectDB()
	if err != nil {
		panic("file khong ton stai")
	}

	db.AutoMigrate(
		&models.Users{},
		&models.PersonalAccessToken{},
	)
}

func main() {
	route.RegisterRoute()
}

// nodemon --exec go run main.go --signal SIGTERM
