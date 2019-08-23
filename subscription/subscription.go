package subscription

import (
	"log"

	"github.com/weriKK/microservice"
)

// Service implements the business logic for creating, retrieving and updating
// monitoring event subscriptions
type Service struct {
	subscriptionDB microservice.SubscriptionStorage
}

// NewService ...
func NewService(storage microservice.SubscriptionStorage) microservice.SubscriptionManager {
	return &Service{
		subscriptionDB: storage,
	}
}

// Subscribe ...
func (s *Service) Subscribe(afID string, mes microservice.MonitoringEventSubscription) (referenceID int, err error) {
	log.Printf("subscription.Subscribe: %#+v, %#+v\n", afID, mes)
	return s.subscriptionDB.Save(afID, mes)
}

// GetSubscription ...
func (s *Service) GetSubscription(afID string, referenceID int) (microservice.MonitoringEventSubscription, error) {
	log.Printf("subscription.GetSubscription: %#+v, %#+v\n", afID, referenceID)
	return s.subscriptionDB.Get(afID, referenceID)
}

// GetSubscriptions ...
func (s *Service) GetSubscriptions(afID string) (map[int]microservice.MonitoringEventSubscription, error) {
	log.Printf("subscription.GetSubscriptions: %#+v\n", afID)
	return s.subscriptionDB.GetAll(afID)
}

// Unsubscribe ...
func (s *Service) Unsubscribe(afID string, referenceID int) error {
	log.Printf("subscription.Unsubscribe: %#+v, %#+v\n", afID, referenceID)
	return s.subscriptionDB.Delete(afID, referenceID)
}
