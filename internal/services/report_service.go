package services

import (
	"workhub/config"
	"workhub/internal/dto"
	"workhub/internal/models"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) GetDashboardReport() (
	*dto.DashboardReportResponse,
	error,
) {

	var totalUsers int64
	var totalCompanies int64
	var approvedCompanies int64
	var pendingCompanies int64
	var totalJobs int64
	var totalApplications int64

	config.DB.
		Model(&models.User{}).
		Count(&totalUsers)

	config.DB.
		Model(&models.Company{}).
		Count(&totalCompanies)

	config.DB.
		Model(&models.Company{}).
		Where(
			"verification_status = ?",
			"approved",
		).
		Count(&approvedCompanies)

	config.DB.
		Model(&models.Company{}).
		Where(
			"verification_status = ?",
			"pending",
		).
		Count(&pendingCompanies)

	config.DB.
		Model(&models.Job{}).
		Count(&totalJobs)

	config.DB.
		Model(&models.Application{}).
		Count(&totalApplications)

	report :=
		dto.DashboardReportResponse{
			TotalUsers:        totalUsers,
			TotalCompanies:    totalCompanies,
			ApprovedCompanies: approvedCompanies,
			PendingCompanies:  pendingCompanies,
			TotalJobs:         totalJobs,
			TotalApplications: totalApplications,
		}

	return &report, nil
}
