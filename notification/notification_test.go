package notification

import (
	"testing"

	"github.com/weriKK/microservice"
	"github.com/weriKK/microservice/mock"
)

func TestSave(t *testing.T) {
	storage := &mock.MysqlStorageService{}
	sender := &mock.NotificationSenderService{}

	service := NewService(storage, sender)

	// Mock storage.Get() call
	storage.GetFn = func(afID string, referenceID int) (microservice.MonitoringEventSubscription, error) {
		if afID == "" {
			t.Fatalf("afID cannot be empty: %q", afID)
		}

		if referenceID <= 0 {
			t.Fatalf("referenceID is expected to be larger than ZERO, but received %d", referenceID)
		}

		return microservice.MonitoringEventSubscription{
			EventType:                      "TestEvent",
			NotificationDestinationAddress: "http://127.0.0.1:9/unreachable",
			MaxReports:                     42,
		}, nil
	}

	// Mock sender.Send() call
	sender.SendFn = func(dest string, mer microservice.MonitoringEventReport) error {
		if dest == "" {
			t.Fatalf("destination address cannot be empty: %q", dest)
		}

		return nil
	}

	// Invoke our service
	mer := microservice.MonitoringEventReport{}
	err := service.NotifySubscriber("test_afID", 1337, mer)
	if err != nil {
		t.Fatalf("service.NotifySubscriber was expected to succeed, but it failed: %v", err)
	}

	if !storage.GetInvoked {
		t.Fatalf("storage.Get() was not invoked")
	}

	if !sender.SendInvoked {
		t.Fatalf("sender.Send() was not invoked")
	}
}
