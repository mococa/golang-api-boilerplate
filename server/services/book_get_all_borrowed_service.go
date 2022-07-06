package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
)

/* -------------- Services -------------- */
func (service *Service) GetAllBorrowedBooksService(book *dtos.BorrowedBooksDTO) (*[]models.BorrowedBook, error) {
	invalid_error := dtos.IsValidBorrowedBooksDTO(book)

	if invalid_error != nil {
		return nil, invalid_error
	}

	borrowed_books, err := service.book_repository.BorrowedBooks(*book)

	return borrowed_books, err
}
