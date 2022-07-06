package services

import (
	"github.com/mococa/golang-api-boilerplate/server/repositories"
)

type Service struct {
	user_repository *repositories.UserRepository
	book_repository *repositories.BookRepository
}
