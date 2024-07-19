package service

import (
	model "go-api/app/domain/models"
	"go-api/app/infra"
)

type SubscriptionService struct {
	repository infra.SubscriptionRepository
}

func NewSubscriptionService(repository infra.SubscriptionRepository) SubscriptionService {
	return SubscriptionService{
		repository,
	}
}

func (ss *SubscriptionService) Subscribe(userId uint, eventId uint) error {
	subscription := model.Subscription{
		EventId: eventId,
		UserId:  userId,
	}

	err := ss.repository.Subscribe(&subscription)

	return err
}

func (ss *SubscriptionService) Unsubscribe(userId uint, eventId uint) error {
	subscription := model.Subscription{
		EventId: eventId,
		UserId:  userId,
	}

	err := ss.repository.Unsubscribe(&subscription)

	return err
}
