package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Producer -
type Producer struct {
	Writer *kafka.Writer
	dialer *kafka.Dialer
	leader string
}

// NewProducer -
func NewProducer(address, topic string) (*Producer, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	p := &Producer{
		leader: address,
		dialer: dialer,
	}

	return p, nil
}

// SendMessage - Send a message to the nominated topic.
func (p *Producer) SendMessage(ctx context.Context, key []byte, value []byte, topic string) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(p.leader),
		Balancer: &kafka.LeastBytes{},
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("Closing writer generated error: %v", err)
		}
	}()

	return writer.WriteMessages(ctx, kafka.Message{
		Topic:  topic,
		Offset: 0,
		Key:    key,
		Value:  value,
	})
}
