package data

import (
	"log"
	"proxy-golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=proxy-golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Enable auto migrations
	err = db.AutoMigrate(&models.Server{}, &models.RedirectServer{})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
