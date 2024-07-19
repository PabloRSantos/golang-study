package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	USER_ROLE  Role = "USER"
	ADMIN_ROLE Role = "ADMIN"
)

type User struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	Email     string         `json:"email" gorm:"size:255;unique;not null"`
	Phone     string         `json:"phoneNumber" gorm:"size:12"`
	Password  string         `json:"-" gorm:"not null"`
	Role      Role           `json:"role" gorm:"not null;default:USER"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Events []Event `json:"events" gorm:"many2many:subscriptions"`
}
