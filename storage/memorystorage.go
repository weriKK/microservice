package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/weriKK/microservice"
)

// InMemoryStorageService ...
type InMemoryStorageService struct {
	sync.RWMutex
	data    map[string]map[int]microservice.MonitoringEventSubscription
	nextIdx map[string]int
}

// NewInMemoryStorageService ...
func NewInMemoryStorageService() microservice.SubscriptionStorage {

	return &InMemoryStorageService{
		data:    map[string]map[int]microservice.MonitoringEventSubscription{},
		nextIdx: map[string]int{},
	}
}

// Save ...
func (imss *InMemoryStorageService) Save(afID string, mes microservice.MonitoringEventSubscription) (referenceID int, err error) {
	log.Printf("storage.Save: %+v, %+v", afID, mes)
	imss.Lock()
	defer imss.Unlock()

	referenceID = imss.nextIdx[afID]
	if 0 == referenceID {
		imss.data[afID] = map[int]microservice.MonitoringEventSubscription{}
	}
	imss.data[afID][referenceID] = mes
	imss.nextIdx[afID]++

	return referenceID, nil
}

// Get ...
func (imss *InMemoryStorageService) Get(afID string, referenceID int) (microservice.MonitoringEventSubscription, error) {
	log.Printf("storage.Get: %+v, %+v", afID, referenceID)
	imss.RLock()
	defer imss.RUnlock()

	if _, ok := imss.data[afID][referenceID]; !ok {
		return microservice.MonitoringEventSubscription{}, fmt.Errorf("subscription not found")
	}
	return imss.data[afID][referenceID], nil
}

// GetAll ...
func (imss *InMemoryStorageService) GetAll(afID string) (map[int]microservice.MonitoringEventSubscription, error) {
	log.Printf("storage.GetAll: %+v", afID)
	imss.RLock()
	defer imss.RUnlock()
	return imss.data[afID], nil
}

// Delete ...
func (imss *InMemoryStorageService) Delete(afID string, referenceID int) error {
	log.Printf("storage.Delete: %+v, %+v", afID, referenceID)
	imss.Lock()
	defer imss.Unlock()
	delete(imss.data[afID], referenceID)
	return nil
}

// IncrementSentReportCount ...
func (imss *InMemoryStorageService) IncrementSentReportCount(afID string, referenceID int) error {
	log.Printf("storage.IncrementSentReportCount: %+v, %+v", afID, referenceID)
	imss.Lock()
	defer imss.Unlock()
	tmp := imss.data[afID][referenceID]
	tmp.MaxReports++
	imss.data[afID][referenceID] = tmp
	return nil
}

// Close ...
func (imss *InMemoryStorageService) Close() error {
	return nil
}
