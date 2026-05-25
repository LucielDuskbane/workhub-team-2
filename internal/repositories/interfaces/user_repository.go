package interfaces

import "workhub/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}
