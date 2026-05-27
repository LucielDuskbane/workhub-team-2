package interfaces

import "workhub/internal/models"

type ApplicationRepository interface {
	Create(application *models.Application) error
	FindByJobAndUser(jobID uint, userID uint) (*models.Application, error)
	GetMyApplications(userID uint) ([]models.Application, error)
	GetByID(id uint) (*models.Application, error)
	GetByJobID(jobID uint) ([]models.Application, error)
	Update(application *models.Application) error
}
