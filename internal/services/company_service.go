package services

import (
	"errors"
	"workhub/internal/dto"
	"workhub/internal/models"
	"workhub/internal/repositories"
	"workhub/internal/repositories/interfaces"
)

type CompanyService struct {
	companyRepo interfaces.CompanyRepository
}

func NewCompanyService() *CompanyService {
	return &CompanyService{
		companyRepo: repositories.NewCompanyRepository(),
	}
}

func (s *CompanyService) CreateCompany(userID uint, req dto.CreateCompanyRequest) error {
	existingCompany, _ := s.companyRepo.GetByUserID(userID)
	if existingCompany.ID != 0 {
		return errors.New("company already exists")
	}

	company := models.Company{
		UserID:             userID,
		CompanyName:        req.CompanyName,
		Description:        req.Description,
		CompanyType:        req.CompanyType,
		VerificationStatus: "pending",
	}
	return s.companyRepo.Create(&company)
}

func (s *CompanyService) GetMyCompany(userID uint) (*models.Company, error) {
	return s.companyRepo.GetByUserID(userID)
}

func (s *CompanyService) UpdateCompany(userID uint, req dto.UpdateCompanyRequest) error {
	company, err := s.companyRepo.GetByUserID(userID)
	if err != nil {
		return errors.New("company not found")
	}

	company.CompanyName = req.CompanyName
	company.Description = req.Description
	company.CompanyType = req.CompanyType

	return s.companyRepo.Update(company)
}

func (s *CompanyService) GetPendingCompanies() ([]models.Company, error) {
	return s.companyRepo.GetPendingCompanies()
}

func (s *CompanyService) ApproveCompany(id uint) error {
	company, err := s.companyRepo.GetByID(id)
	if err != nil {
		return errors.New("company not found")
	}

	company.VerificationStatus = "approved"
	return s.companyRepo.Update(company)
}

func (s *CompanyService) RejectCompany(id uint) error {
	company, err := s.companyRepo.GetByID(id)
	if err != nil {
		return errors.New("company not found")
	}

	company.VerificationStatus = "rejected"
	return s.companyRepo.Update(company)
}
