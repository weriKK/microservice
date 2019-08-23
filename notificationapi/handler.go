package notificationapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/weriKK/microservice"
)

// NotificationServiceHandler ...
type NotificationServiceHandler struct {
	*mux.Router
	service microservice.MonitoringEventReporter
}

// NewNotificationServiceHandler ...
func NewNotificationServiceHandler(service microservice.MonitoringEventReporter) *NotificationServiceHandler {
	h := &NotificationServiceHandler{
		Router:  mux.NewRouter(),
		service: service,
	}

	h.Methods("POST").Path("/notify/{afID}/{referenceID}").HandlerFunc(h.notificationHandler)

	return h
}

func (ns *NotificationServiceHandler) notificationHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("notificationapi: %s - %s", r.Method, r.RequestURI)

	afID := mux.Vars(r)["afID"]
	referenceID, err := strconv.Atoi(mux.Vars(r)["referenceID"])
	if err != nil {
		log.Println("notificationapi: Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var mer microservice.MonitoringEventReport
	err = json.NewDecoder(r.Body).Decode(&mer)
	if err != nil {
		log.Println("notificationapi: Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("notificationapi: body: %#+v", mer)

	err = ns.service.NotifySubscriber(afID, referenceID, mer)
	if err != nil {
		log.Println("notificationapi: Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
