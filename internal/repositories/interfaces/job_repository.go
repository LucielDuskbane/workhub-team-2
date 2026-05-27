package interfaces

import "workhub/internal/models"

type JobRepository interface {
	Create(job *models.Job) error
	FindAll() ([]models.Job, error)
	FindByID(id uint) (*models.Job, error)
	FindByCompanyID(companyID uint) ([]models.Job, error)
	Update(job *models.Job) error
	Delete(id uint) error
}
