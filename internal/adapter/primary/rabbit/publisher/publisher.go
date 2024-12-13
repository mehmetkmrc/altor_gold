package publisher

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/mehmetkmrc/ator_gold/internal/adapter/primary/rabbit"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/logger"
	"time"
	"github.com/google/uuid"
	"go.uber.org/zap"
)


const (
	_exchangeKind		= "direct"
	_exchangeDurable	= true
	_exchangeAutoDelete = false
	_exchangeInternal	= false
	_exchangeNoWait		= false
	_queueDurable		= true
	_queueAutoDelete	= false
	_queueExclusive		= false
	_queueNoWait 		= false
	_publishMandatory	= false
	_publishImmediate   = false
	_ExchangeName       = "x-exchange"
	_BindingKey         = "x-routing-key"
	_MessageTypeName    = "x"
)

type (
	QueueNames []string
	Publisher struct{
		Queues 				 QueueNames
		ExchangeName, BindingKey string
		MessageTypeName 	 string
		QueueArg 			 amqp.Table
		AmqpConn 			 *amqp.Connection
	}
)

var _ EventPublisher = (*Publisher)(nil)

func NewPublisher(AmqpConn *amqp.Connection, config *Publisher)(EventPublisher, error) {
	pub := &Publisher{
		AmqpConn: AmqpConn,
		ExchangeName: config.ExchangeName,
		Queues: config.Queues,
		QueueArg: config.QueueArg,
		BindingKey: config.BindingKey,
		MessageTypeName: config.MessageTypeName,
	}

	return pub, nil
}

func (p *Publisher) Configure(opts ...Option) EventPublisher{
	for _, opt := range opts {
		opt(p)
	}

	return p 
}

func (p *Publisher) PublishEvents(ctx context.Context, events []any) error {
	for _, e := range events {
		b, err := json.Marshal(e)
		if err != nil {
			return errors.Wrap(err, "Publisher-json.Marshal")
		}

		err = p.Publish(ctx, b, rabbitmq.ContentTypeJSON)
		if err != nil {
			return errors.Wrap(err, "Publisher-pub.Publish")
		}
	}

	return nil
}

func (p *Publisher) Publish(ctx context.Context, body []byte, contentType string) error {
	ch, err := p.AmqpConn.Channel()
	defer func() {
		err := ch.Close()
		if err != nil {
			zap.S().Error("Error while closing channel: ", err)
		}
	}()

	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}

	zap.S().Info("publishing message to exchange: ", p.ExchangeName)
	if err := ch.PublishWithContext(
		ctx,
		p.ExchangeName,
		p.BindingKey,
		_publishMandatory,
		_publishImmediate,
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         body,
			Type:         p.MessageTypeName,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}

	return nil
}

func (p *Publisher) SetupExchangeAndQueues() error {
	ch, err := p.AmqpConn.Channel()
	defer func() {
		err := ch.Close()
		if err != nil {
			zap.S().Error("Error while closing channel: ", err)
		}
	}()

	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	err = ch.ExchangeDeclare(
		p.ExchangeName,
		_exchangeKind,
		_exchangeDurable,
		_exchangeAutoDelete,
		_exchangeInternal,
		_exchangeNoWait,
		p.QueueArg,
	)
	if err != nil {
		return errors.Wrap(err, "Error ch.ExchangeDeclare")
	}

	for _, queueName := range p.Queues {
		_, err := ch.QueueDeclare(
			queueName,
			_queueDurable,
			_queueAutoDelete,
			_queueExclusive,
			_queueNoWait,
			p.QueueArg,
		)
		if err != nil {
			return errors.Wrap(err, "Error ch.QueueDeclare")
		}

		err = ch.QueueBind(
			queueName,
			p.BindingKey,
			p.ExchangeName,
			_queueNoWait,
			p.QueueArg,
		)
		if err != nil {
			return errors.Wrap(err, "Error ch.QueueBind")
		}

		zap.S().Info("Queue:\t", queueName, "\tbound to the exchange:\t", p.ExchangeName)
	}

	return nil
}

func (p *Publisher) GetExchangeName() string {
	return p.ExchangeName
}

func (p *Publisher) Close() {
	isClosed := p.AmqpConn.IsClosed()

	if !isClosed {
		err := p.AmqpConn.Close()
		if err != nil {
			zap.S().Errorf("Error while disconnecting from RabbitMQ: %s", err)
		}
	}

	logger.ForceLog("Connection to RabbitMQ closed successfully")
}