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

func (h *CompanyHandler) CreateCompany(
	c *gin.Context,
) {
	var req dto.CreateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	userID := c.MustGet(
		"user_id",
	).(uint)

	err := h.companyService.
		CreateCompany(
			userID,
			req,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"Company created successfully",
		nil,
	)
}

func (h *CompanyHandler) GetMyCompany(
	c *gin.Context,
) {
	userID := c.MustGet(
		"user_id",
	).(uint)

	company, err :=
		h.companyService.
			GetMyCompany(userID)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"Company not found",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		company,
	)
}

func (h *CompanyHandler) UpdateCompany(
	c *gin.Context,
) {
	var req dto.UpdateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	userID := c.MustGet(
		"user_id",
	).(uint)

	err := h.companyService.
		UpdateCompany(
			userID,
			req,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Company updated successfully",
		nil,
	)
}

func (h *CompanyHandler) GetPendingCompanies(
	c *gin.Context,
) {
	companies, err :=
		h.companyService.
			GetPendingCompanies()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed get companies",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		companies,
	)
}

func (h *CompanyHandler) ApproveCompany(
	c *gin.Context,
) {
	idParam := c.Param("id")

	id, err :=
		strconv.ParseUint(
			idParam,
			10,
			32,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid ID",
		)
		return
	}

	err = h.companyService.
		ApproveCompany(
			uint(id),
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Company approved",
		nil,
	)
}

func (h *CompanyHandler) RejectCompany(
	c *gin.Context,
) {
	idParam := c.Param("id")

	id, err :=
		strconv.ParseUint(
			idParam,
			10,
			32,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid ID",
		)
		return
	}

	err = h.companyService.
		RejectCompany(
			uint(id),
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Company rejected",
		nil,
	)
}
