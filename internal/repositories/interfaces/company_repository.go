package interfaces

import "workhub/internal/models"

type CompanyRepository interface {
	Create(company *models.Company) error
	GetByUserID(userID uint) (*models.Company, error)
	GetPendingCompanies() ([]models.Company, error)
	Update(company *models.Company) error
	GetByID(id uint) (*models.Company, error)
}
