package model

import (
	"time"
)

const TableNameEmergencyContact = "emergency_contacts"

// EmergencyContact mapped from table <emergency_contacts>
type EmergencyContact struct {
	ID           int     `gorm:"column:id;primaryKey" json:"id"`
	UserID       int     `gorm:"column:user_id;not null" json:"user_id"`
	Code         string    `gorm:"column:code;not null;default:gen_random_uuid()" json:"code"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	Phone        string    `gorm:"column:phone;not null" json:"phone"`
	Relationship string    `gorm:"column:relationship;not null" json:"relationship"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName EmergencyContact's table name
func (*EmergencyContact) TableName() string {
	return TableNameEmergencyContact
}
