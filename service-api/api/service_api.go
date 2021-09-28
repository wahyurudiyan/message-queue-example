package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/wahyurudiyan/message-queue-example/service-api/model"
	"github.com/wahyurudiyan/message-queue-example/service-api/publisher"
)

type api struct {
	queueName string
	publisher publisher.Publisher
}

type API interface {
	SendNotification(c echo.Context) error
}

func NewService(queueName string, publisher publisher.Publisher) API {
	return &api{
		queueName: queueName,
		publisher: publisher,
	}
}

func (a *api) SendNotification(c echo.Context) error {
	var notifObj model.Notification

	err := c.Bind(&notifObj)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if len(notifObj.Receivers) == 0 {
		return c.JSON(http.StatusBadRequest, errors.New("receiver is nil, please insert at least one"))
	} else if notifObj.Title == "" {
		return c.JSON(http.StatusBadRequest, errors.New("title is empty, please insert your notification title"))
	}

	a.publisher.PublishMessage(a.queueName, notifObj)

	return c.JSON(http.StatusOK, map[string]string{"message": "published"})
}
