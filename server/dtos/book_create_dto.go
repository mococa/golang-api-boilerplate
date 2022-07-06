package dtos

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type CreateBookDTO struct {
	Name        string `json:"name" valid:"required"`
	Author      string `json:"author" valid:"required"`
	ReleaseYear int    `json:"release_year" valid:"required"`
}

func IsValidCreateBookDTO(dto *CreateBookDTO) error {
	result, _ := govalidator.ValidateStruct(dto)

	// TODO: Properly handle validation errors
	if result == false {
		return errors.New("Invalid body")
	}

	return nil
}
