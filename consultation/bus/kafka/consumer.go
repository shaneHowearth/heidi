package kafka

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
	dialer *kafka.Dialer
	Topic  string
}

// Start from the beginning of the queue that this client knows about.
const startOffset = -2

var (
	BadAddress = errors.New("poorly formed address")
	BadTopic   = errors.New("poorly formed topic")
)

func NewConsumer(address, topic string) (*Consumer, error) {
	if address == "" {
		return nil, BadAddress
	}
	if topic == "" {
		return nil, BadTopic
	}
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	c := Consumer{
		dialer: dialer,
		Topic:  topic,
	}
	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     c.Topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		// MaxWait:   time.Millisecond * 10,
		Dialer: c.dialer,
	})

	if err := c.Reader.SetOffset(startOffset); err != nil {
		log.Print("Unable to set offset in reader")
		return nil, err
	}

	return &c, nil

}

func (c *Consumer) Read(count int) ([][]byte, error) {
	data := make([][]byte, count)

	for i := 0; i < count; i++ {
		message, err := c.Reader.ReadMessage(context.Background())
		if err != nil {
			return nil, err
		}

		data[i] = message.Value
	}

	return data, nil
}
