package repositories

import (
	"workhub/config"
	"workhub/internal/models"
)

type ApplicationRepository struct{}

func NewApplicationRepository() *ApplicationRepository {
	return &ApplicationRepository{}
}

func (r *ApplicationRepository) Create(
	application *models.Application,
) error {
	return config.DB.
		Create(application).
		Error
}

func (r *ApplicationRepository) FindByJobAndUser(
	jobID uint,
	userID uint,
) (*models.Application, error) {

	var application models.Application

	err := config.DB.
		Where(
			"job_id = ? AND jobseeker_id = ?",
			jobID,
			userID,
		).
		First(&application).Error

	return &application, err
}

func (r *ApplicationRepository) GetMyApplications(
	userID uint,
) ([]models.Application, error) {

	var applications []models.Application

	err := config.DB.
		Where(
			"jobseeker_id = ?",
			userID,
		).
		Find(&applications).Error

	return applications, err
}

func (r *ApplicationRepository) GetByID(
	id uint,
) (*models.Application, error) {

	var application models.Application

	err := config.DB.
		First(&application, id).
		Error

	return &application, err
}

func (r *ApplicationRepository) GetByJobID(
	jobID uint,
) ([]models.Application, error) {

	var applications []models.Application

	err := config.DB.
		Where(
			"job_id = ?",
			jobID,
		).
		Find(&applications).Error

	return applications, err
}

func (r *ApplicationRepository) Update(
	application *models.Application,
) error {
	return config.DB.
		Save(application).
		Error
}
