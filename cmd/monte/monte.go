package main

import (
	"log"

	"github.com/weriKK/microservice/http"
	"github.com/weriKK/microservice/storage"
	"github.com/weriKK/microservice/subscription"
	"github.com/weriKK/microservice/subscriptionapi"
)

func main() {
	log.Println("Subscription Microservice")

	//inMemoryStorage := storage.NewInMemoryStorageService()
	sqlStorage, err := storage.NewMysqlStorageService()
	if err != nil {
		panic(err)
	}
	subscriptionService := subscription.NewService(sqlStorage)

	apiHandler := subscriptionapi.NewMonitoringEventServiceHandler(subscriptionService)
	apiServer := http.NewServer(":8100", apiHandler)

	log.Fatal(apiServer.ListenAndServe())
}
