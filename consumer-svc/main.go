package main

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/wahyurudiyan/message-queue-example/consumer-svc/consumer"
)

func main() {
	queueName := "myQueue"
	consSvc := consumer.NewConsumer("guest", "guest", "localhost", "5672")

	message, err := consSvc.ConsumeMessage(queueName)
	if err != nil {
		logrus.Errorf("[error] unable to consume: %s", err)
	}

	loop := make(chan bool)
	go func() {
		for m := range message {
			data, _ := json.Marshal(string(m.Body))
			logrus.Printf("DATA : %s", string(data))
		}
	}()

	fmt.Printf("[start] message started to listen, to exit pser CTRL+C\n\n")
	<-loop // this channel will block your code until code interrupt
}
