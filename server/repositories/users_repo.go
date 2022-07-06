package repositories

import (
	"github.com/mococa/golang-api-boilerplate/server/dtos"
	"github.com/mococa/golang-api-boilerplate/server/models"

	"gorm.io/gorm"
)

/* -------------- Types -------------- */
type UserRepository struct {
	Db *gorm.DB
}

/* -------------- Initialization -------------- */
func NewUsersRepo(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

/* -------------- Repository Methods -------------- */
func (repository *UserRepository) CreateUser(user dtos.CreateUserDTO) (*models.User, error) {
	var created_user models.User = models.User{
		Email:    user.Email,
		Password: user.Password,
		Picture:  user.Picture,
		Name:     user.Name,
	}

	err := repository.Db.Create(&created_user).Error

	if err != nil {
		return nil, err
	}

	return &created_user, nil
}

func (repository *UserRepository) ListUsers() (*[]models.User, error) {
	users := []models.User{}

	err := repository.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}
