package mock

import "github.com/weriKK/microservice"

// MysqlStorageService is a mock implementation of the microservice.SubscriptionStorage interface
type MysqlStorageService struct {
	SaveFn      func(afID string, mes microservice.MonitoringEventSubscription) (referenceID int, err error)
	SaveInvoked bool

	GetFn      func(afID string, referenceID int) (microservice.MonitoringEventSubscription, error)
	GetInvoked bool

	GetAllFn      func(afID string) (map[int]microservice.MonitoringEventSubscription, error)
	GetAllInvoked bool

	DeleteFn      func(afID string, referenceID int) error
	DeleteInvoked bool

	IncrementSentReportCountFn      func(afID string, referenceID int) error
	IncrementSentReportCountInvoked bool

	CloseFn      func() error
	CloseInvoked bool
}

// Save invokes the mock implementation and marks the function as invoked
func (m *MysqlStorageService) Save(afID string, mes microservice.MonitoringEventSubscription) (referenceID int, err error) {
	m.SaveInvoked = true
	return m.SaveFn(afID, mes)
}

// Get ...
func (m *MysqlStorageService) Get(afID string, referenceID int) (microservice.MonitoringEventSubscription, error) {
	m.GetInvoked = true
	return m.GetFn(afID, referenceID)
}

// GetAll ...
func (m *MysqlStorageService) GetAll(afID string) (map[int]microservice.MonitoringEventSubscription, error) {
	m.GetAllInvoked = true
	return m.GetAllFn(afID)
}

// Delete ...
func (m *MysqlStorageService) Delete(afID string, referenceID int) error {
	m.DeleteInvoked = true
	return m.DeleteFn(afID, referenceID)
}

// IncrementSentReportCount ...
func (m *MysqlStorageService) IncrementSentReportCount(afID string, referenceID int) error {
	m.IncrementSentReportCountInvoked = true
	return m.IncrementSentReportCountFn(afID, referenceID)
}

// Close ...
func (m *MysqlStorageService) Close() error {
	m.CloseInvoked = true
	return m.CloseFn()
}
