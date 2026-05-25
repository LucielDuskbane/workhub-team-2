package dto

type CreateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Salary      int64  `json:"salary"`
	Location    string `json:"location"`
	JobType     string `json:"job_type"`
}

type UpdateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Salary      int64  `json:"salary"`
	Location    string `json:"location"`
	JobType     string `json:"job_type"`
	Status      string `json:"status"`
}
