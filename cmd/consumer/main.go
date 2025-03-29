package main

import (
	"github.com/sirupsen/logrus"
	"kafkaGolang/internal/handler"
	"kafkaGolang/internal/kafka"
	"os"
	"os/signal"
	"syscall"
)

const (
	topic         = "my-topic"
	consumerGroup = "my-consumer-group"
)

var address = []string{"localhost:9092", "localhost:9093", "localhost:9094"}

func main() {
	h := handler.NewHandler()
	c1, err := kafka.NewConsumer(h, address, topic, consumerGroup, 1)
	if err != nil {
		logrus.Fatal(err)
	}
	c2, err := kafka.NewConsumer(h, address, topic, consumerGroup, 2)
	if err != nil {
		logrus.Fatal(err)
	}
	c3, err := kafka.NewConsumer(h, address, topic, consumerGroup, 3)
	if err != nil {
		logrus.Fatal(err)
	}

	go c1.Start()
	go c2.Start()
	go c3.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	logrus.Fatal(c1.Stop(), c2.Stop(), c3.Stop())
}
