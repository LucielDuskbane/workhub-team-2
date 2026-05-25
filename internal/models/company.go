package models

import "time"

type Company struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	UserID             uint      `gorm:"unique;not null" json:"user_id"`
	CompanyName        string    `gorm:"type:varchar(100);not null" json:"company_name"`
	Description        string    `gorm:"type:text" json:"description"`
	CompanyType        string    `gorm:"type:varchar(50)" json:"company_type"`
	VerificationStatus string    `gorm:"default:pending" json:"verification_status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
