package infra

import (
	"fmt"
	model "go-api/app/domain/models"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	connection *gorm.DB
}

func NewSubscriptionRepository(connection *gorm.DB) SubscriptionRepository {
	return SubscriptionRepository{
		connection,
	}
}

func (sr *SubscriptionRepository) Subscribe(subscription *model.Subscription) error {
	err := sr.connection.Create(subscription).Error

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (sr *SubscriptionRepository) Unsubscribe(subscription *model.Subscription) error {
	err := sr.connection.Delete(subscription).Error

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (sr *SubscriptionRepository) GetByEvent(eventId uint) []model.Subscription {
	var subscriptions []model.Subscription

	sr.connection.Where(&model.Subscription{EventId: eventId}).Find(&subscriptions)

	return subscriptions
}

func (sr *SubscriptionRepository) CountByEvent(eventId uint) int64 {
	var totalSubscriptions int64

	sr.connection.Model(&model.Subscription{}).Where(&model.Subscription{EventId: eventId}).Count(&totalSubscriptions)

	return totalSubscriptions
}

func (sr *SubscriptionRepository) GetByUser(userId uint) []model.Subscription {
	var subscriptions []model.Subscription

	sr.connection.Where(&model.Subscription{UserId: userId}).Find(&subscriptions)

	return subscriptions
}
