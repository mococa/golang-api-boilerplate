package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
	"github.com/mococa/golang-api-boilerplate/server/repositories"
)

/* -------------- Types -------------- */
type UserService interface {
	CreateUserService(user *dtos.CreateUserDTO) (*models.User, error)
	ListUserService() (*[]models.User, error)
}

/* -------------- Initialization -------------- */
func NewUserService(user_repository *repositories.UserRepository) *Service {
	return &Service{user_repository: user_repository}
}
