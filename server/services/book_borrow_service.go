package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
)

/* -------------- Services -------------- */
func (service *Service) BorrowBookService(book *dtos.BorrowBookDTO) (*models.BorrowedBook, error) {
	invalid_error := dtos.IsValidBorrowBookDTO(book)

	if invalid_error != nil {
		return nil, invalid_error
	}

	borrowed_book, err := service.book_repository.BorrowBook(*book)

	return borrowed_book, err
}
