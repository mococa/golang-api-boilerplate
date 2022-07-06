package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
)

/* -------------- Services -------------- */
func (service *Service) CreateBookService(book *dtos.CreateBookDTO) (*models.Book, error) {
	invalid_error := dtos.IsValidCreateBookDTO(book)

	if invalid_error != nil {
		return nil, invalid_error
	}

	created_book, err := service.book_repository.CreateBook(*book)

	return created_book, err
}
