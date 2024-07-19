package infra

import (
	"fmt"
	model "go-api/app/domain/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	connection *gorm.DB
}

func NewEventRepository(connection *gorm.DB) EventRepository {
	return EventRepository{
		connection: connection,
	}
}

func (er *EventRepository) GetEvents() []model.Event {
	var events []model.Event
	er.connection.Find(&events)

	return events
}

func (er *EventRepository) CreateEvent(event *model.Event) error {
	result := er.connection.Create(&event)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (er *EventRepository) GetEventById(id uint) (*model.Event, error) {
	var event model.Event
	err := er.connection.First(&event, id).Error

	return &event, err
}

func (er *EventRepository) GetEventUsers(id uint) ([]model.User, error) {
	var event model.Event
	err := er.connection.Preload("Users").First(&event, id).Error

	return event.Users, err
}
