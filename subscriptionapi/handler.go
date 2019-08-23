package subscriptionapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/weriKK/microservice"	
)

// SubscriptionServiceHandler ...
type SubscriptionServiceHandler struct {
	*mux.Router
	service microservice.SubscriptionManager
}

// NewMonitoringEventServiceHandler ...
func NewMonitoringEventServiceHandler(service microservice.SubscriptionManager) *SubscriptionServiceHandler {
	h := &SubscriptionServiceHandler{
		Router:  mux.NewRouter(),
		service: service,
	}

	h.Methods("POST").Path("/subscription/{afID}").HandlerFunc(h.createSubscriptionHandler)

	return h
}

func (ss *SubscriptionServiceHandler) createSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("subscriptionapi: %s - %s", r.Method, r.RequestURI)

	afID := mux.Vars(r)["afID"]

	var mes microservice.MonitoringEventSubscription
	err := json.NewDecoder(r.Body).Decode(&mes)
	if err != nil {
		log.Println("subscriptionapi: Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("subscriptionapi: body: %#+v", mes)

	referenceID, err := ss.service.Subscribe(afID, mes)
	if err != nil {
		log.Println("subscriptionapi: Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("subscriptionapi: %#+v", referenceID)

	w.WriteHeader(http.StatusOK)
}
