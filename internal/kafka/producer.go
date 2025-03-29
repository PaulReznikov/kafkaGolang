package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"strings"
	"time"
)

const (
	flushTimeout = 5000 //ms
)

var errUnknownType error = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(address []string) (*Producer, error) {

	conf := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("error with new producer: %w", err)
	}
	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message, topic, key string, timeStamp time.Time) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:       &topic,
			Partition:   kafka.PartitionAny,
			Offset:      0,
			Metadata:    nil,
			Error:       nil,
			LeaderEpoch: nil,
		},
		Value:     []byte(message),
		Key:       []byte(key),
		Timestamp: timeStamp,
	}

	kafkaChan := make(chan kafka.Event)

	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return fmt.Errorf("error with delivery of message: %w", err)
	}

	e := <-kafkaChan
	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
