package database

import (
	_ "challenge-api/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "putra"
	password = "test123"
	dbPort   = "5432"
	dbname   = "mini-shop"
	DB       *gorm.DB
)


func DBInit() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	// url := "root:@tcp(localhost:5432)/mini-shop"

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to database :")
	}

	// fmt.Println("connecting to database")
	// db.Debug().AutoMigrate(&models.Item{}, &models.Order{})
	

	DB = db
	return db, nil
}
