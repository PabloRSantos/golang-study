package service

import (
	dto "go-api/app/domain/dtos"
	model "go-api/app/domain/models"
	"go-api/app/infra"
)

type EventService struct {
	repository infra.EventRepository
}

func NewEventService(repository infra.EventRepository) EventService {
	return EventService{
		repository: repository,
	}
}

func (eu *EventService) GetEvents() []model.Event {
	return eu.repository.GetEvents()
}

func (eu *EventService) CreateEvent(payload dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	event := model.Event{
		Name:        payload.Name,
		Description: &payload.Description,
		Date:        payload.Date,
	}

	err := eu.repository.CreateEvent(&event)

	if err != nil {
		return dto.CreateEventResponse{}, err
	}

	response := dto.CreateEventResponse{
		ID: event.ID,
	}
	return response, nil
}

func (eu *EventService) GetEventById(id int) (*model.Event, error) {
	return eu.repository.GetEventById(id)
}
