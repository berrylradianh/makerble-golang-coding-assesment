package model

import (
	"time"
)

const TableNameMedicalRecord = "medical_records"

// MedicalRecord mapped from table <medical_records>
type MedicalRecord struct {
	ID               int     `gorm:"column:id;primaryKey" json:"id"`
	Code             string    `gorm:"column:code;not null" json:"code"`
	UserID           int     `gorm:"column:user_id;not null" json:"user_id"`
	DoctorID         int     `gorm:"column:doctor_id;not null" json:"doctor_id"`
	CurrentComplaint string    `gorm:"column:current_complaint;not null" json:"current_complaint"`
	DiseaseHistory   *string   `gorm:"column:disease_history" json:"disease_history"`
	MedicineAllergy  *string   `gorm:"column:medicine_allergy" json:"medicine_allergy"`
	MedicationTaken  *string   `gorm:"column:medication_taken" json:"medication_taken"`
	IsEverSurgery    bool      `gorm:"column:is_ever_surgery;not null" json:"is_ever_surgery"`
	AssignedAt       time.Time `gorm:"column:assigned_at;not null" json:"assigned_at"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName MedicalRecord's table name
func (*MedicalRecord) TableName() string {
	return TableNameMedicalRecord
}
