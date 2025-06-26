package model

import (
	"time"
)

const TableNameUserDoctor = "user_doctors"

// UserDoctor mapped from table <user_doctors>
type UserDoctor struct {
	ID                        int     `gorm:"column:id;primaryKey" json:"id"`
	Code                      string    `gorm:"column:code;not null;default:gen_random_uuid()" json:"code"`
	UserID                    int     `gorm:"column:user_id;not null" json:"user_id"`
	EmploymentStatusID        int     `gorm:"column:employment_status_id;not null" json:"employment_status_id"`
	DepartmentID              int     `gorm:"column:department_id;not null" json:"department_id"`
	ProfessionalDegree        string    `gorm:"column:professional_degree;not null" json:"professional_degree"`
	Citizenship               string    `gorm:"column:citizenship;not null" json:"citizenship"`
	Language                  string    `gorm:"column:language;not null" json:"language"`
	GmcNumber                 string    `gorm:"column:gmc_number;not null" json:"gmc_number"`
	Specialization            *string   `gorm:"column:specialization" json:"specialization"`
	SubSpecialization         *string   `gorm:"column:sub_specialization" json:"sub_specialization"`
	LastEducation             string    `gorm:"column:last_education;not null" json:"last_education"`
	YearGraduated             string    `gorm:"column:year_graduated;not null" json:"year_graduated"`
	ProfessionalCertification *string   `gorm:"column:professional_certification" json:"professional_certification"`
	EmployeeNumber            string    `gorm:"column:employee_number;not null" json:"employee_number"`
	IsActive                  bool      `gorm:"column:is_active;not null" json:"is_active"`
	CreatedAt                 time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName UserDoctor's table name
func (*UserDoctor) TableName() string {
	return TableNameUserDoctor
}
