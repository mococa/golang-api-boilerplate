package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	DB_URI := "postgres://root:secret@localhost/bookdb?sslmode=disable"

	db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	runMigrations(db)

	return db
}

func GetDB() *gorm.DB {
	return db
}
