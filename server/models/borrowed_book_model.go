package models

import (
	uuid "github.com/satori/go.uuid"
)

type BorrowedBook struct { // table name: borrowed_books
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`

	Book   *Book     `json:"book" gorm:"foreignKey:BookID"`
	BookID uuid.UUID `json:"book_id"`

	User   *User     `json:"user" gorm:"foreignKey:UserID"`
	UserID uuid.UUID `json:"user_id" gorm:"index:idx_user_book"`
}
