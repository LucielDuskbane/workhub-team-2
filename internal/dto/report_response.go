package dto

type DashboardReportResponse struct {
	TotalUsers        int64 `json:"total_users"`
	TotalCompanies    int64 `json:"total_companies"`
	ApprovedCompanies int64 `json:"approved_companies"`
	PendingCompanies  int64 `json:"pending_companies"`
	TotalJobs         int64 `json:"total_jobs"`
	TotalApplications int64 `json:"total_applications"`
}
