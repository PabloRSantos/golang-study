package service

import (
	dto "go-api/app/domain/dtos"
	model "go-api/app/domain/models"
	"go-api/app/infra"
)

type EventService struct {
	eventRepository        infra.EventRepository
	subscriptionRepository infra.SubscriptionRepository
}

func NewEventService(eventRepository infra.EventRepository, subscriptionRepository infra.SubscriptionRepository) EventService {
	return EventService{
		eventRepository,
		subscriptionRepository,
	}
}

func (eu *EventService) GetEvents() []model.Event {
	return eu.eventRepository.GetEvents()
}

func (eu *EventService) CreateEvent(payload dto.CreateEventRequest) (dto.CreateEventResponse, error) {
	event := model.Event{
		Name:        payload.Name,
		Description: &payload.Description,
		Date:        payload.Date,
	}

	err := eu.eventRepository.CreateEvent(&event)

	if err != nil {
		return dto.CreateEventResponse{}, err
	}

	response := dto.CreateEventResponse{
		ID: event.ID,
	}
	return response, nil
}

func (eu *EventService) GetEventById(id uint) (dto.GetEventResponse, error) {
	event, err := eu.eventRepository.GetEventById(id)
	if err != nil {
		return dto.GetEventResponse{}, err
	}

	subscriptions := eu.subscriptionRepository.CountByEvent(id)

	response := dto.GetEventResponse{
		Event:         *event,
		Subscriptions: subscriptions,
	}

	return response, nil
}

func (eu *EventService) GetEventUsers(id uint) ([]model.User, error) {
	event, err := eu.eventRepository.GetEventById(id)
	if err != nil {
		return []model.User{}, err
	}

	users, err := eu.eventRepository.GetEventUsers(event.ID)
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
