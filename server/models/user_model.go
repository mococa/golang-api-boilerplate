package models

import (
	uuid "github.com/satori/go.uuid"
)

type User struct { // table name: users
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`

	Email    string `json:"email" gorm:"unique;unique_index"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}
