package database

import (
	"challenge-api/pkg/models"
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


func DBInit() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	// url := "root:@tcp(localhost:5432)/mini-shop"

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to database :")
	}

	// fmt.Println("connecting to database")
	db.Debug().AutoMigrate(&models.User{})
	

	// DB = db
	// return db, nil
}

func GetDB() *gorm.DB {
	return DB
}
