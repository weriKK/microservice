package notification

import (
	"log"

	"github.com/weriKK/microservice"
)

// Service implements the business logic for passing monitoring event reports
// to a subscriber
type Service struct {
	subscriptionDB     microservice.SubscriptionStorage
	notificationSender microservice.NotificationSender
}

// NewService ...
func NewService(storage microservice.SubscriptionStorage, sender microservice.NotificationSender) microservice.MonitoringEventReporter {
	return &Service{
		subscriptionDB:     storage,
		notificationSender: sender,
	}
}

// NotifySubscriber ...
func (s *Service) NotifySubscriber(afID string, referenceID int, mer microservice.MonitoringEventReport) error {
	log.Printf("notification.NotifySubscriber: %+v, %+v, %+v", afID, referenceID, mer)

	sub, err := s.subscriptionDB.Get(afID, referenceID)
	if err != nil {
		return err
	}

	return s.notificationSender.Send(sub.NotificationDestinationAddress, mer)
}
