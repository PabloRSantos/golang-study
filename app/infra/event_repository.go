package infra

import (
	"fmt"
	domain "go-api/app/domain/models"

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

func (er *EventRepository) GetEvents() []domain.Event {
	var events []domain.Event
	er.connection.Find(&events)

	return events
}

func (er *EventRepository) CreateEvent(event *domain.Event) error {
	result := er.connection.Create(&event)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (er *EventRepository) GetEventById(id int) (*domain.Event, error) {
	var event domain.Event
	err := er.connection.First(&event, id).Error

	return &event, err
}
