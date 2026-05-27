package handlers

import (
	"net/http"
	"strconv"
	"workhub/internal/dto"
	"workhub/internal/services"
	"workhub/internal/utils"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func NewCompanyHandler() *CompanyHandler {
	return &CompanyHandler{
		companyService: services.NewCompanyService(),
	}
}

// Create Company godoc
// @Summary Create company
// @Description Employer create company profile
// @Tags Companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCompanyRequest true "Create Company"
// @Success 201 {object} map[string]interface{}
// @Router /companies [post]
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req dto.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.MustGet("user_id").(uint)
	if err := h.companyService.CreateCompany(userID, req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Company created successfully", nil)
}

// Get My Company godoc
// @Summary Get my company
// @Description Employer get own company
// @Tags Companies
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /companies/me [get]
func (h *CompanyHandler) GetMyCompany(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	company, err := h.companyService.GetMyCompany(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Company not found")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Success", company)
}

// Update Company godoc
// @Summary Update company
// @Description Employer update company
// @Tags Companies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateCompanyRequest true "Update Company"
// @Success 200 {object} map[string]interface{}
// @Router /companies [put]
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	var req dto.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.MustGet("user_id").(uint)
	if err := h.companyService.UpdateCompany(userID, req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Company updated successfully", nil)
}

// Get Pending Companies godoc
// @Summary Get pending companies
// @Description Admin get pending companies
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /admin/companies/pending [get]
func (h *CompanyHandler) GetPendingCompanies(c *gin.Context) {
	companies, err := h.companyService.GetPendingCompanies()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed get companies")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Success", companies)
}

// Approve Company godoc
// @Summary Approve company
// @Description Admin approve company
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Company ID"
// @Success 200 {object} map[string]interface{}
// @Router /admin/companies/{id}/approve [patch]
func (h *CompanyHandler) ApproveCompany(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.companyService.ApproveCompany(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Company approved", nil)
}

// Reject Company godoc
// @Summary Reject company
// @Description Admin reject company
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "Company ID"
// @Success 200 {object} map[string]interface{}
// @Router /admin/companies/{id}/reject [patch]
func (h *CompanyHandler) RejectCompany(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.companyService.RejectCompany(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Company rejected", nil)
}
