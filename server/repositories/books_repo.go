package repositories

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"

	"gorm.io/gorm"
)

/* -------------- Types -------------- */
type BookRepository struct {
	Db *gorm.DB
}

/* -------------- Initialization -------------- */
func NewBooksRepo(db *gorm.DB) BookRepository {
	return BookRepository{db}
}

/* -------------- Repository Methods -------------- */
func (repository *BookRepository) CreateBook(book dtos.CreateBookDTO) (*models.Book, error) {
	var created_book models.Book = models.Book{
		Name:        book.Name,
		Author:      book.Author,
		ReleaseYear: book.ReleaseYear,
	}

	err := repository.Db.Create(&created_book).Error

	if err != nil {
		return nil, err
	}

	return &created_book, nil
}

func (repository *BookRepository) BorrowBook(book dtos.BorrowBookDTO) (*models.BorrowedBook, error) {
	var borrowed_book models.BorrowedBook = models.BorrowedBook{
		BookID: book.BookID,
		UserID: book.UserID,
	}

	err := repository.Db.Create(&borrowed_book).Error

	if err != nil {
		return nil, err
	}

	return &borrowed_book, nil
}

func (repository *BookRepository) BorrowedBooks(borrowed dtos.BorrowedBooksDTO) (*[]models.BorrowedBook, error) {
	var borrowed_books []models.BorrowedBook

	err := repository.Db.Joins("Book").Joins("User").Find(&borrowed_books, "user_id =?", borrowed.UserID).Error

	if err != nil {
		return nil, err
	}

	return &borrowed_books, nil
}
