package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int     `gorm:"column:id;primaryKey" json:"id"`
	Code           string    `gorm:"column:code;not null;default:gen_random_uuid()" json:"code"`
	RoleID         int     `gorm:"column:role_id;not null" json:"role_id"`
	IdentityNumber string    `gorm:"column:identity_number;not null" json:"identity_number"`
	Email          string    `gorm:"column:email;not null" json:"email"`
	Password       string    `gorm:"column:password;not null" json:"password"`
	Name           string    `gorm:"column:name;not null" json:"name"`
	Phone          string    `gorm:"column:phone;not null" json:"phone"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth;not null" json:"date_of_birth"`
	Address        string    `gorm:"column:address;not null" json:"address"`
	Gender         string    `gorm:"column:gender;not null" json:"gender"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
