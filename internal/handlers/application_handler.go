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

func (h *ApplicationHandler) ApplyJob(
	c *gin.Context,
) {
	idParam := c.Param("id")

	jobID, err :=
		strconv.ParseUint(
			idParam,
			10,
			32,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid job ID",
		)
		return
	}

	userID := c.MustGet(
		"user_id",
	).(uint)

	err = h.applicationService.
		ApplyJob(
			uint(jobID),
			userID,
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
		"Application submitted",
		nil,
	)
}

func (h *ApplicationHandler) GetMyApplications(
	c *gin.Context,
) {
	userID := c.MustGet(
		"user_id",
	).(uint)

	applications, err :=
		h.applicationService.
			GetMyApplications(
				userID,
			)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed get applications",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		applications,
	)
}

func (h *ApplicationHandler) GetJobApplications(
	c *gin.Context,
) {
	idParam := c.Param("id")

	jobID, err :=
		strconv.ParseUint(
			idParam,
			10,
			32,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"Invalid job ID",
		)
		return
	}

	applications, err :=
		h.applicationService.
			GetJobApplications(
				uint(jobID),
			)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed get applications",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		applications,
	)
}

func (h *ApplicationHandler) UpdateApplicationStatus(
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

	var req dto.UpdateApplicationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	err = h.applicationService.
		UpdateApplicationStatus(
			uint(id),
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
		"Application updated",
		nil,
	)
}
