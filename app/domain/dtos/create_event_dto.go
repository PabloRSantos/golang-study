package dto

import "time"

type CreateEventRequest struct {
	Name        string    `form:"name" binding:"required"`
	Description string    `form:"description"`
	Date        time.Time `form:"date" binding:"required"`
}

type CreateEventResponse struct {
	ID uint `json:"id"`
}
