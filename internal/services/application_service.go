package services

import (
	"errors"
	"time"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
)

type ApplicationService struct {
	applicationRepo *repositories.ApplicationRepository
	jobRepo         *repositories.JobRepository
	companyRepo     *repositories.CompanyRepository
}

func NewApplicationService() *ApplicationService {
	return &ApplicationService{
		applicationRepo: repositories.NewApplicationRepository(),

		jobRepo: repositories.NewJobRepository(),

		companyRepo: repositories.NewCompanyRepository(),
	}
}

func (s *ApplicationService) ApplyJob(
	jobID uint,
	userID uint,
) error {

	_, err :=
		s.jobRepo.FindByID(jobID)

	if err != nil {
		return errors.New(
			"job not found",
		)
	}

	existingApplication, _ :=
		s.applicationRepo.
			FindByJobAndUser(
				jobID,
				userID,
			)

	if existingApplication.ID != 0 {
		return errors.New(
			"already applied",
		)
	}

	application := models.Application{
		JobID:       jobID,
		JobseekerID: userID,
		Status:      "pending",
		AppliedAt:   time.Now(),
	}

	return s.applicationRepo.
		Create(&application)
}

func (s *ApplicationService) GetMyApplications(
	userID uint,
) ([]models.Application, error) {

	return s.applicationRepo.
		GetMyApplications(userID)
}

func (s *ApplicationService) GetJobApplications(
	jobID uint,
) ([]models.Application, error) {

	return s.applicationRepo.
		GetByJobID(jobID)
}

func (s *ApplicationService) UpdateApplicationStatus(
	id uint,
	req dto.UpdateApplicationRequest,
) error {

	application, err :=
		s.applicationRepo.GetByID(id)

	if err != nil {
		return errors.New(
			"application not found",
		)
	}

	application.Status =
		req.Status

	return s.applicationRepo.
		Update(application)
}
