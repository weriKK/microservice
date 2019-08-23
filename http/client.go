package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/weriKK/microservice"
)

// NotificationSenderService implements an HTTP client for sending MonitoringEventReports to subscribers
type NotificationSenderService struct {
}

// NewNotificationSenderService ...
func NewNotificationSenderService() microservice.NotificationSender {
	return &NotificationSenderService{}
}

// Send ...
func (n *NotificationSenderService) Send(dest string, mer microservice.MonitoringEventReport) error {
	log.Printf("http.Send: %#+v, %#+v\n", dest, mer)

	body, err := json.Marshal(mer)
	if err != nil {
		return err
	}

	resp, err := http.Post(dest, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error response for notification request: %v", err)
	}

	return nil
}
