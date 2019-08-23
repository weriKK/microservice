package microservice

// MonitoringEventSubscription represents an AF's subscription to a monitoring event
type MonitoringEventSubscription struct {
	Self                           string `json:"self,omitempty"`
	EventType                      string `json:"eventType"`
	NotificationDestinationAddress string `json:"notificationDestinationAddr"`
	MaxReports                     int    `json:"maxReports"`
}

// SubscriptionManager provides a way for consumers to subscribe, unsubscribe and list subscriptions for monitoring event
type SubscriptionManager interface {
	Subscribe(afID string, mes MonitoringEventSubscription) (referenceID int, err error)
	GetSubscription(afID string, referenceID int) (MonitoringEventSubscription, error)
	GetSubscriptions(afID string) (map[int]MonitoringEventSubscription, error)
	Unsubscribe(afID string, referenceID int) error
}

// MonitoringEventReport represents a report to an AF about a monitoring event
type MonitoringEventReport struct {
	EventType      string `json:"eventType"`
	Report         string `json:"report"`
	SubscriptionID string `json:"subscriptionId,omitempty"`
}

// MonitoringEventReporter provides a way to send notifications about monitoring events
// to subscribers
type MonitoringEventReporter interface {
	NotifySubscriber(afID string, referenceID int, mer MonitoringEventReport) error
}

// SubscriptionStorage provides persistance storage for MonitoringEventSubscription data
type SubscriptionStorage interface {
	Save(afID string, mes MonitoringEventSubscription) (referenceID int, err error)
	Get(afID string, referenceID int) (MonitoringEventSubscription, error)
	GetAll(afID string) (map[int]MonitoringEventSubscription, error)
	Delete(afID string, referenceID int) error
	IncrementSentReportCount(afID string, referenceID int) error
	Close() error
}

// NotificationSender provides a way for consumers to send MonitoringEventReports to a destination address
type NotificationSender interface {
	Send(dest string, mer MonitoringEventReport) error
}

// HTTPAPIServer provides an HTTP server for handling REST API requests
type HTTPAPIServer interface {
	ListenAndServe() error
	Shutdown() error
}
