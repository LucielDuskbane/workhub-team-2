package repositories

import (
	"workhub/config"
	"workhub/internal/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	return &user, err
}
