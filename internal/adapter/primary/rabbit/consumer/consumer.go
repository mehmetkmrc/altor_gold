package consumer

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	_queueDurable     = true
	_queueAutoDelete  = false
	_queueExclusive   = false
	_queueNoWait      = false
	_prefetchCount    = 2000
	_prefetchSize     = 0
	_prefetchGlobal   = false
	_consumeAutoAck   = false
	_consumeExclusive = false
	_consumeNoLocal   = false
	_consumeNoWait    = false
)

type Consumer struct {
	AmqpConn               *amqp.Connection
	QueueName, ConsumerTag string
	QueueArg               amqp.Table
	WorkerPoolSize         int
}

var _ EventConsumer = (*Consumer)(nil)

func NewConsumer(AmqpConn *amqp.Connection, config *Consumer) (EventConsumer, error) {
	sub := &Consumer{
		AmqpConn:       AmqpConn,
		QueueName:      config.QueueName,
		ConsumerTag:    config.ConsumerTag,
		QueueArg:       config.QueueArg,
		WorkerPoolSize: config.WorkerPoolSize,
	}

	return sub, nil
}

func (c *Consumer) Configure(opts ...Option) EventConsumer {
	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Consumer) StartConsumer(ctx context.Context, fn Worker) error {
	consumerWaitGroup := new(sync.WaitGroup)
	ch, err := c.createChannel()
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()

	deliveries, err := ch.Consume(
		c.QueueName,
		c.ConsumerTag,
		_consumeAutoAck,
		_consumeExclusive,
		_consumeNoLocal,
		_consumeNoWait,
		c.QueueArg,
	)

	consumerQueue := fmt.Sprintf("Consumer - %s", c.QueueName)
	if err != nil {
		return errors.Wrap(err, consumerQueue)
	}

	for i := 0; i < c.WorkerPoolSize; i++ {
		consumerWaitGroup.Add(1)
		go func(ctx context.Context, deliveries <-chan amqp.Delivery) {
			defer consumerWaitGroup.Done()
			fn(ctx, deliveries)
		}(ctx, deliveries)
	}

	<-ctx.Done()
	err = ch.Cancel(c.ConsumerTag, false)
	if err != nil {
		zap.S().Error("Error while cancel ch.Cancel: ", err)
	}

	consumerWaitGroup.Wait()
	zap.S().Info("Consumer stopped")
	return ctx.Err()
}

func (c *Consumer) createChannel() (*amqp.Channel, error) {
	ch, err := c.AmqpConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "Error AmqpConn.Channel")
	}

	zap.S().Info("declaring queue:", c.QueueName)
	_, err = ch.QueueDeclare(
		c.QueueName,
		_queueDurable,
		_queueAutoDelete,
		_queueExclusive,
		_queueNoWait,
		c.QueueArg,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
	}
	zap.S().Info("queue bound to exchange, starting to consume from queue. consumer_tag: ", c.ConsumerTag)
	err = ch.Qos(
		_prefetchCount,
		_prefetchSize,
		_prefetchGlobal,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error ch.Qos")
	}

	return ch, nil
}
