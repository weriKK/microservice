package mock

import "github.com/weriKK/microservice"

// NotificationSenderService is a mock implementation of the microservice.NotificationSender interface
type NotificationSenderService struct {
	SendFn      func(dest string, mer microservice.MonitoringEventReport) error
	SendInvoked bool
}

// Send ...
func (n *NotificationSenderService) Send(dest string, mer microservice.MonitoringEventReport) error {
	n.SendInvoked = true
	return n.SendFn(dest, mer)
}
