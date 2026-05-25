package models

import "time"

type Job struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CompanyID   uint      `gorm:"not null" json:"company_id"`
	Title       string    `gorm:"type:varchar(100);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"type:varchar(50)" json:"category"`
	Salary      int64     `json:"salary"`
	Location    string    `gorm:"type:varchar(100)" json:"location"`
	JobType     string    `gorm:"type:varchar(50)" json:"job_type"`
	Status      string    `gorm:"default:open" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
