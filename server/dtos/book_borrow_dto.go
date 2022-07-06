package dtos

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type BorrowBookDTO struct {
	UserID uuid.UUID `json:"user_id" valid:"required"`
	BookID uuid.UUID `json:"book_id" valid:"required"`
}

func IsValidBorrowBookDTO(dto *BorrowBookDTO) error {
	result, _ := govalidator.ValidateStruct(dto)

	// TODO: Properly handle validation errors
	if result == false {
		return errors.New("Invalid body")
	}

	return nil
}
