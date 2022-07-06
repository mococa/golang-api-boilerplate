package services

import (
	"github.com/mococa/golang-api-boilerplate/server/models"
)

/* -------------- Services -------------- */
func (service *Service) ListUserService() (*[]models.User, error) {

	users, err := service.user_repository.ListUsers()

	return users, err
}
