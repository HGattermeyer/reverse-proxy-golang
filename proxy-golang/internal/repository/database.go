package repository

import (
	"fmt"
	"log"
	"os"
	"proxy-golang/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {

	// Retry loop to establish the database connection due to docker compose
	var db *gorm.DB
	var err error
	connectionString := "host=postgres user=admin password=admin dbname=proxy-golang port=5432 sslmode=disable "
	for retries := 0; retries < 10; retries++ {

		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to the database: %v\n", err)
		time.Sleep(5 * time.Second)
	}

	if db == nil {
		fmt.Println("Failed to establish a connection to the database after retries.")
		os.Exit(1)
	}

	// Your application logic here
	fmt.Println("Connected to the database.")
	// Enable auto migrations
	err = db.AutoMigrate(&models.Server{}, &models.RedirectServer{})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
