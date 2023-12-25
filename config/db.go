package config

import (
	"fmt"

	"github.com/narasimha-1511/zolo-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	// Do something here.
	// db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"))

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Borrowed{})
	
	DB = db

}