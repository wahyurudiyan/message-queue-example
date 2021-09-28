package publisher

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type publisher struct {
	amqpClient *amqp.Connection
}

type Publisher interface {
	PublishMessage(queueName string, message interface{}) error
}

func NewPublisher(username, password, baseUrl, port string) Publisher {
	if port == "" {
		port = "5672"
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, baseUrl, port)
	amqpClient, err := amqp.Dial(url)
	if err != nil {
		logrus.Panicf("[panic] unable to define amqp: %s", err)
	}

	return &publisher{amqpClient}
}

func (p *publisher) PublishMessage(queueName string, message interface{}) error {
	ch, err := p.amqpClient.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	m, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        m,
	})
}
