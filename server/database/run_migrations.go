package database

import (
	"github.com/mococa/golang-api-boilerplate/server/models"

	"gorm.io/gorm"
)

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		/* -------------- User -------------- */
		&models.User{},

		/* -------------- Book -------------- */
		&models.Book{}, &models.BorrowedBook{},
	)
}
