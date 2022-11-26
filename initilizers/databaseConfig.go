package initializers

import (
	"log"
	"os"

	"github.com/nidhey27/ticket-booking-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to DB")
	} else {
		DB.AutoMigrate(&models.Admin{})
		log.Printf("Connected to DB")

	}
}