package repositories

import (
	"workhub/config"
	"workhub/internal/models"
)

type JobRepository struct{}

func NewJobRepository() *JobRepository {
	return &JobRepository{}
}

func (r *JobRepository) Create(
	job *models.Job,
) error {
	return config.DB.Create(job).Error
}

func (r *JobRepository) FindAll() (
	[]models.Job,
	error,
) {
	var jobs []models.Job

	err :=
		config.DB.Find(&jobs).Error

	return jobs, err
}

func (r *JobRepository) FindByID(
	id uint,
) (*models.Job, error) {

	var job models.Job

	err := config.DB.
		First(&job, id).Error

	return &job, err
}

func (r *JobRepository) FindByCompanyID(
	companyID uint,
) ([]models.Job, error) {

	var jobs []models.Job

	err := config.DB.
		Where(
			"company_id = ?",
			companyID,
		).
		Find(&jobs).Error

	return jobs, err
}

func (r *JobRepository) Update(
	job *models.Job,
) error {
	return config.DB.Save(job).Error
}

func (r *JobRepository) Delete(
	id uint,
) error {
	return config.DB.
		Delete(
			&models.Job{},
			id,
		).
		Error
}
