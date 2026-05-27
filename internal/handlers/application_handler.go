package handlers

import (
	"net/http"
	"strconv"
	"workhub/internal/dto"
	"workhub/internal/services"
	"workhub/internal/utils"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	applicationService *services.ApplicationService
}

func NewApplicationHandler() *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: services.NewApplicationService(),
	}
}

// Apply Job godoc
// @Summary Apply to job
// @Description Jobseeker apply job
// @Tags Applications
// @Produce json
// @Security BearerAuth
// @Param id path int true "Job ID"
// @Success 201 {object} map[string]interface{}
// @Router /jobs/{id}/apply [post]
func (h *ApplicationHandler) ApplyJob(c *gin.Context) {
	idParam := c.Param("id")

	jobID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid job ID")
		return
	}

	userID := c.MustGet("user_id").(uint)

	if err := h.applicationService.ApplyJob(uint(jobID), userID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusCreated, "Application submitted", nil)
}

// Get My Applications godoc
// @Summary Get my applications
// @Description Jobseeker get own applications
// @Tags Applications
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /applications/me [get]
func (h *ApplicationHandler) GetMyApplications(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	applications, err := h.applicationService.GetMyApplications(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed get applications")
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Success", applications)
}

// Get Job Applications godoc
// @Summary Get job applications
// @Description Employer get applicants
// @Tags Applications
// @Produce json
// @Security BearerAuth
// @Param id path int true "Job ID"
// @Success 200 {object} map[string]interface{}
// @Router /jobs/{id}/applications [get]
func (h *ApplicationHandler) GetJobApplications(c *gin.Context) {
	idParam := c.Param("id")

	jobID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid job ID")
		return
	}

	userID := c.MustGet("user_id").(uint)

	applications, err := h.applicationService.GetJobApplications(uint(jobID), userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Success", applications)
}

// Update Application Status godoc
// @Summary Update application status
// @Description Employer accept/reject applicant
// @Tags Applications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Application ID"
// @Param request body dto.UpdateApplicationRequest true "Update Status"
// @Success 200 {object} map[string]interface{}
// @Router /applications/{id} [patch]
func (h *ApplicationHandler) UpdateApplicationStatus(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req dto.UpdateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.MustGet("user_id").(uint)

	if err := h.applicationService.UpdateApplicationStatus(uint(id), userID, req); err != nil {
		utils.ErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Application updated", nil)
}
