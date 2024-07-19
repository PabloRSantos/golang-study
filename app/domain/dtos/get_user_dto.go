package dto

import (
	model "go-api/app/domain/models"
	"time"
)

type eventResponse struct {
	ID   uint      `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type GetUserResponse struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Role      model.Role      `json:"role"`
	Phone     string          `json:"phoneNumber"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Events    []eventResponse `json:"events"`
}

func NewGetUserResponse(user model.User) GetUserResponse {
	var eventsResponse []eventResponse

	for _, event := range user.Events {
		formattedEvent := eventResponse{
			ID:   event.ID,
			Name: event.Name,
			Date: event.Date,
		}

		eventsResponse = append(eventsResponse, formattedEvent)
	}

	return GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Events:    eventsResponse,
	}
}
