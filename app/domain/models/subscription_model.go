package model

import (
	"time"
)

type Subscription struct {
	EventId   uint      `json:"eventId" gorm:"primaryKey"`
	UserId    uint      `json:"userId" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`

	Event Event `json:"event"`
	User  User  `json:"user"`
}
