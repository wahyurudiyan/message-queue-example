package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/wahyurudiyan/message-queue-example/service-api/api"
	"github.com/wahyurudiyan/message-queue-example/service-api/publisher"
)

func main() {
	RESTPort := ":8080"
	pubSvc := publisher.NewPublisher("guest", "guest", "localhost", "5672")
	handler := api.NewService("myQueue", pubSvc)

	ec := echo.New()
	router := ec.Group("/api/v1")
	router.POST("/send/notification", handler.SendNotification)

	if err := ec.Start(RESTPort); err != nil {
		logrus.Fatalf("[fail] server unable to start: %s", err)
	}
}
