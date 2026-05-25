package models

import "time"

type Application struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	JobID       uint      `gorm:"not null" json:"job_id"`
	JobseekerID uint      `gorm:"not null" json:"jobseeker_id"`
	Status      string    `gorm:"default:pending" json:"status"`
	AppliedAt   time.Time `json:"applied_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
