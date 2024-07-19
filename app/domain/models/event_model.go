package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	Description *string        `json:"description"`
	Date        time.Time      `json:"date" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`

	Users []User `gorm:"many2many:subscriptions"`
}
