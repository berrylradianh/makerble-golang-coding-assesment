package model

import (
	"time"
)

const TableNameEmploymentStatus = "employment_statuses"

// EmploymentStatus mapped from table <employment_statuses>
type EmploymentStatus struct {
	ID        int     `gorm:"column:id;primaryKey" json:"id"`
	Code      string    `gorm:"column:code;not null;default:gen_random_uuid()" json:"code"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Slug      string    `gorm:"column:slug;not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName EmploymentStatus's table name
func (*EmploymentStatus) TableName() string {
	return TableNameEmploymentStatus
}
