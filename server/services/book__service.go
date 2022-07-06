package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
	"github.com/mococa/golang-api-boilerplate/server/repositories"
)

/* -------------- Types -------------- */
type BookService interface {
	CreateBookService(book *dtos.CreateBookDTO) (*models.Book, error)
	BorrowBookService(book *dtos.BorrowBookDTO) (*models.BorrowedBook, error)
	GetAllBorrowedBooksService(book *dtos.BorrowedBooksDTO) (*[]models.BorrowedBook, error)
}

/* -------------- Initialization -------------- */
func NewBookService(book_repository *repositories.BookRepository) *Service {
	return &Service{book_repository: book_repository}
}
