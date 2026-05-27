package services

import (
	"errors"
	"log"
	"time"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
	"workhub/internal/repositories/interfaces"
	"workhub/internal/utils"
)

type ApplicationService struct {
	applicationRepo interfaces.ApplicationRepository
	jobRepo         interfaces.JobRepository
	companyRepo     interfaces.CompanyRepository
	userRepo        interfaces.UserRepository
}

func NewApplicationService() *ApplicationService {
	return &ApplicationService{
		applicationRepo: repositories.NewApplicationRepository(),
		jobRepo:         repositories.NewJobRepository(),
		companyRepo:     repositories.NewCompanyRepository(),
		userRepo:        repositories.NewUserRepository(),
	}
}

func (s *ApplicationService) ApplyJob(jobID uint, userID uint) error {
	if _, err := s.jobRepo.FindByID(jobID); err != nil {
		return errors.New("job not found")
	}

	existingApplication, _ := s.applicationRepo.FindByJobAndUser(jobID, userID)
	if existingApplication.ID != 0 {
		return errors.New("already applied")
	}

	application := models.Application{
		JobID:       jobID,
		JobseekerID: userID,
		Status:      "pending",
		AppliedAt:   time.Now(),
	}
	return s.applicationRepo.Create(&application)
}

func (s *ApplicationService) GetMyApplications(userID uint) ([]models.Application, error) {
	return s.applicationRepo.GetMyApplications(userID)
}

func (s *ApplicationService) GetJobApplications(jobID uint, userID uint) ([]models.Application, error) {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("company not found")
	}

	job, err := s.jobRepo.FindByID(jobID)
	if err != nil {
		return nil, errors.New("job not found")
	}

	if job.CompanyID != company.ID {
		return nil, errors.New("forbidden access")
	}
	return s.applicationRepo.GetByJobID(jobID)
}

func (s *ApplicationService) UpdateApplicationStatus(id uint, userID uint, req dto.UpdateApplicationRequest) error {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return errors.New("company not found")
	}

	application, err := s.applicationRepo.GetByID(id)
	if err != nil {
		return errors.New("application not found")
	}

	job, err := s.jobRepo.FindByID(application.JobID)
	if err != nil {
		return errors.New("job not found")
	}

	if job.CompanyID != company.ID {
		return errors.New("forbidden access")
	}

	application.Status = req.Status
	if err := s.applicationRepo.Update(application); err != nil {
		return err
	}

	// Send email if accepted
	if req.Status == "accepted" {
		user, err := s.userRepo.FindByID(application.JobseekerID)
		if err != nil {
			return err
		}

		subject := "Job Application Accepted"
		content := "<h1>Congratulations!</h1><p>Your application has been accepted.</p>"
		go func(email string, name string) {
			if err := utils.SendEmail(email, name, subject, content); err != nil {
				log.Printf("failed to send email to %s: %v", email, err)
			}
		}(user.Email, user.Name)
	}
	return nil
}
