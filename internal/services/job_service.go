package services

import (
	"errors"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
)

type JobService struct {
	jobRepo     *repositories.JobRepository
	companyRepo *repositories.CompanyRepository
}

func NewJobService() *JobService {
	return &JobService{
		jobRepo: repositories.NewJobRepository(),

		companyRepo: repositories.NewCompanyRepository(),
	}
}

func (s *JobService) CreateJob(
	userID uint,
	req dto.CreateJobRequest,
) error {

	company, err :=
		s.companyRepo.
			GetByUserID(userID)

	if err != nil {
		return errors.New(
			"company not found",
		)
	}

	if company.VerificationStatus !=
		"approved" {

		return errors.New(
			"company not approved",
		)
	}

	job := models.Job{
		CompanyID:   company.ID,
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Salary:      req.Salary,
		Location:    req.Location,
		JobType:     req.JobType,
		Status:      "open",
	}

	return s.jobRepo.Create(
		&job,
	)
}

func (s *JobService) GetAllJobs() (
	[]models.Job,
	error,
) {
	return s.jobRepo.FindAll()
}

func (s *JobService) GetJobByID(
	id uint,
) (*models.Job, error) {
	return s.jobRepo.FindByID(id)
}
