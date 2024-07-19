package dto

import model "go-api/app/domain/models"

type GetEventResponse struct {
	Event         model.Event `json:"event"`
	Subscriptions int64       `json:"subscriptions"`
}
