package services

import (
	"errors"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
	"workhub/internal/repositories/interfaces"
)

type JobService struct {
	jobRepo     interfaces.JobRepository
	companyRepo interfaces.CompanyRepository
}

func NewJobService() *JobService {
	return &JobService{
		jobRepo:     repositories.NewJobRepository(),
		companyRepo: repositories.NewCompanyRepository(),
	}
}

func (s *JobService) CreateJob(userID uint, req dto.CreateJobRequest) error {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return errors.New("company not found")
	}

	if company.VerificationStatus != "approved" {
		return errors.New("company not approved")
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
	return s.jobRepo.Create(&job)
}

func (s *JobService) GetAllJobs() ([]models.Job, error) {
	return s.jobRepo.FindAll()
}

func (s *JobService) GetJobByID(id uint) (*models.Job, error) {
	return s.jobRepo.FindByID(id)
}

func (s *JobService) GetMyJobs(userID uint) ([]models.Job, error) {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("company not found")
	}

	jobs, err := s.jobRepo.FindByCompanyID(company.ID)
	return jobs, err
}

func (s *JobService) UpdateJob(userID uint, jobID uint, req dto.UpdateJobRequest) error {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return errors.New("company not found")
	}

	job, err := s.jobRepo.FindByID(jobID)
	if err != nil {
		return errors.New("job not found")
	}

	if job.CompanyID != company.ID {
		return errors.New("forbidden access")
	}

	job.Title = req.Title
	job.Description = req.Description
	job.Category = req.Category
	job.Salary = req.Salary
	job.Location = req.Location
	job.JobType = req.JobType
	job.Status = req.Status

	return s.jobRepo.Update(job)
}

func (s *JobService) DeleteJob(userID uint, jobID uint) error {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return errors.New("company not found")
	}

	job, err := s.jobRepo.FindByID(jobID)
	if err != nil {
		return errors.New("job not found")
	}

	if job.CompanyID != company.ID {
		return errors.New("forbidden access")
	}
	return s.jobRepo.Delete(jobID)
}
