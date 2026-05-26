package services

import (
	"errors"
	"time"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
	"workhub/internal/utils"
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
	userID uint,
) ([]models.Application, error) {

	company, err :=
		s.companyRepo.
			GetByUserID(userID)

	if err != nil {
		return nil,
			errors.New(
				"company not found",
			)
	}

	job, err :=
		s.jobRepo.
			FindByID(jobID)

	if err != nil {
		return nil,
			errors.New(
				"job not found",
			)
	}

	if job.CompanyID !=
		company.ID {

		return nil,
			errors.New(
				"forbidden access",
			)
	}

	return s.applicationRepo.
		GetByJobID(jobID)
}

func (s *ApplicationService) UpdateApplicationStatus(
	id uint,
	userID uint,
	req dto.UpdateApplicationRequest,
) error {

	company, err :=
		s.companyRepo.
			GetByUserID(userID)

	if err != nil {
		return errors.New(
			"company not found",
		)
	}

	application, err :=
		s.applicationRepo.
			GetByID(id)

	if err != nil {
		return errors.New(
			"application not found",
		)
	}

	job, err :=
		s.jobRepo.
			FindByID(
				application.JobID,
			)

	if err != nil {
		return errors.New(
			"job not found",
		)
	}

	if job.CompanyID !=
		company.ID {

		return errors.New(
			"forbidden access",
		)
	}

	application.Status =
		req.Status

	err = s.applicationRepo.
		Update(application)

	if err != nil {
		return err
	}

	// Send email if accepted
	if req.Status ==
		"accepted" {

		userRepo :=
			repositories.
				NewUserRepository()

		user, err :=
			userRepo.FindByID(
				application.
					JobseekerID,
			)

		if err != nil {
			return err
		}

		err =
			utils.SendEmail(
				user.Email,
				user.Name,
				"Job Application Accepted",
				"<h1>Congratulations!</h1><p>Your application has been accepted.</p>",
			)

		if err != nil {
			return err
		}
	}

	return nil
}
