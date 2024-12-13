package consumer

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Worker func(ctx context.Context, messages <-chan amqp.Delivery)

type EventConsumer interface {
	Configure(...Option) EventConsumer
	StartConsumer(ctx context.Context, fn Worker) error
}
