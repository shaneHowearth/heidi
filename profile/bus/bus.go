package bus

import "context"

type Carrier interface {
	SendMessage(ctx context.Context, key []byte, value []byte, topic string) error
}
