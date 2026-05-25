package handlers

import (
	"net/http"
	"strconv"
	"workhub/internal/dto"
	"workhub/internal/services"
	"workhub/internal/utils"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService *services.JobService
}

func NewJobHandler() *JobHandler {
	return &JobHandler{
		jobService: services.NewJobService(),
	}
}

func (h *JobHandler) CreateJob(
	c *gin.Context,
) {
	var req dto.CreateJobRequest

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

	err := h.jobService.
		CreateJob(
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
		"Job created successfully",
		nil,
	)
}

func (h *JobHandler) GetAllJobs(
	c *gin.Context,
) {
	jobs, err :=
		h.jobService.
			GetAllJobs()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed get jobs",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		jobs,
	)
}

func (h *JobHandler) GetJobByID(
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

	job, err :=
		h.jobService.
			GetJobByID(
				uint(id),
			)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"Job not found",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		job,
	)
}
