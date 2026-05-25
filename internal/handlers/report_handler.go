package handlers

import (
	"net/http"
	"workhub/internal/services"
	"workhub/internal/utils"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService *services.ReportService
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{
		reportService: services.NewReportService(),
	}
}

func (h *ReportHandler) GetDashboardReport(
	c *gin.Context,
) {

	report, err :=
		h.reportService.
			GetDashboardReport()

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed get report",
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Success",
		report,
	)
}
