package consumer

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type consumer struct {
	amqpClient *amqp.Connection
}

type Consumer interface {
	ConsumeMessage(queueName string) (<-chan amqp.Delivery, error)
}

func NewConsumer(username, password, baseUrl, port string) Consumer {
	if port == "" {
		port = "5672"
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, baseUrl, port)
	amqpClient, err := amqp.Dial(url)
	if err != nil {
		logrus.Panicf("[panic] unable to define amqp: %s", err)
	}

	return &consumer{amqpClient: amqpClient}
}

func (c *consumer) ConsumeMessage(queueName string) (<-chan amqp.Delivery, error) {
	ch, err := c.amqpClient.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	m, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return m, nil
}
