package dto

type CreateCompanyRequest struct {
	CompanyName string `json:"company_name"`
	Description string `json:"description"`
	CompanyType string `json:"company_type"`
}

type UpdateCompanyRequest struct {
	CompanyName string `json:"company_name"`
	Description string `json:"description"`
	CompanyType string `json:"company_type"`
}
