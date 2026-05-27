package repositories

import (
	"workhub/config"
	"workhub/internal/models"
)

type CompanyRepository struct{}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return config.DB.Create(company).Error
}

func (r *CompanyRepository) GetByUserID(userID uint) (*models.Company, error) {
	var company models.Company
	err := config.DB.Where("user_id = ?", userID).First(&company).Error
	return &company, err
}

func (r *CompanyRepository) GetPendingCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := config.DB.Where("verification_status = ?", "pending").Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return config.DB.Save(company).Error
}

func (r *CompanyRepository) GetByID(id uint) (*models.Company, error) {
	var company models.Company
	err := config.DB.First(&company, id).Error
	return &company, err
}
