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

// Create Job godoc
// @Summary Create job
// @Description Employer create new job
// @Tags Jobs
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateJobRequest true "Create Job"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /jobs [post]
func (h *JobHandler) CreateJob(
	c *gin.Context,
) {

	var req dto.CreateJobRequest

	if err :=
		c.ShouldBindJSON(
			&req,
		); err != nil {

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

// Get All Jobs godoc
// @Summary Get all jobs
// @Description Public get all jobs
// @Tags Jobs
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /jobs [get]
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

// Get Job By ID godoc
// @Summary Get job detail
// @Description Get job by ID
// @Tags Jobs
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /jobs/{id} [get]
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

// Get My Jobs godoc
// @Summary Get my jobs
// @Description Employer get own jobs
// @Tags Jobs
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /jobs/my [get]
func (h *JobHandler) GetMyJobs(
	c *gin.Context,
) {

	userID := c.MustGet(
		"user_id",
	).(uint)

	jobs, err :=
		h.jobService.
			GetMyJobs(
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
		http.StatusOK,
		"Success",
		jobs,
	)
}

// Update Job godoc
// @Summary Update job
// @Description Employer update own job
// @Tags Jobs
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Job ID"
// @Param request body dto.UpdateJobRequest true "Update Job"
// @Success 200 {object} map[string]interface{}
// @Router /jobs/{id} [put]
func (h *JobHandler) UpdateJob(
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
			"Invalid ID",
		)
		return
	}

	var req dto.UpdateJobRequest

	if err :=
		c.ShouldBindJSON(
			&req,
		); err != nil {

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

	err = h.jobService.
		UpdateJob(
			userID,
			uint(jobID),
			req,
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusForbidden,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Job updated successfully",
		nil,
	)
}

// Delete Job godoc
// @Summary Delete job
// @Description Employer delete own job
// @Tags Jobs
// @Produce json
// @Security BearerAuth
// @Param id path int true "Job ID"
// @Success 200 {object} map[string]interface{}
// @Router /jobs/{id} [delete]
func (h *JobHandler) DeleteJob(
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
			"Invalid ID",
		)
		return
	}

	userID := c.MustGet(
		"user_id",
	).(uint)

	err = h.jobService.
		DeleteJob(
			userID,
			uint(jobID),
		)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusForbidden,
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Job deleted successfully",
		nil,
	)
}
