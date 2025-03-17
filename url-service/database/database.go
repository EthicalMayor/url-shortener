package database

import (
	"log"
	"url-shortener/url-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:mdmdndshdiwr82482h2kdnIop[0.,;[sj@tcp(localhost:3306)/url?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to the database successfully")

	db.AutoMigrate(&models.URL{})
	DB = db
}
