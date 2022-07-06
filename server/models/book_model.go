package models

import (
	uuid "github.com/satori/go.uuid"
)

type Book struct { // table name: books
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name        string `json:"name" gorm:"index:idx_book,unique"`
	Author      string `json:"author" gorm:"index:idx_book,unique"`
	ReleaseYear int    `json:"release_year"`
}
