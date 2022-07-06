package services

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"
)

/* -------------- Services -------------- */
func (service *Service) CreateUserService(user *dtos.CreateUserDTO) (*models.User, error) {
	invalid_error := dtos.IsValidCreateUserDTO(user)

	if invalid_error != nil {
		return nil, invalid_error
	}

	created_user, err := service.user_repository.CreateUser(*user)

	return created_user, err
}
