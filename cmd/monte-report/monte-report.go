package main

import (
	"log"

	"github.com/weriKK/microservice/http"
	"github.com/weriKK/microservice/notification"
	"github.com/weriKK/microservice/notificationapi"
	"github.com/weriKK/microservice/storage"
)

func main() {
	log.Println("Notification Microservice")

	//inMemoryStorage := storage.NewInMemoryStorageService()
	sqlStorage, err := storage.NewMysqlStorageService()
	if err != nil {
		panic(err)
	}

	httpSender := http.NewNotificationSenderService()
	notificationService := notification.NewService(sqlStorage, httpSender)
	apiHandler := notificationapi.NewNotificationServiceHandler(notificationService)
	apiServer := http.NewServer(":8101", apiHandler)

	log.Fatal(apiServer.ListenAndServe())
}
