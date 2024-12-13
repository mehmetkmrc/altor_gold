package rabbitmq

import (
	"errors"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"time"

	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	_retryTimes     = 5
	_backOffSeconds = 2
)

var (
	ErrCannotConnectRabbitMQ = errors.New("cannot connect to rabbit")
	ContentTypeJSON          = "application/json"
	ContentTypeText          = "text/plain"
)

func NewRabbitMQConn(cfg *config.Container) (*amqp.Connection, error) {
	var (
		amqpConn *amqp.Connection
		counts   int64
	)

	zap.S().Info("Connection string: ", cfg.RabbitMQ.URL)
	for {
		connection, err := amqp.Dial(cfg.RabbitMQ.URL)
		if err != nil {
			zap.S().Error("failed to connect to RabbitMq...", err, cfg.RabbitMQ.URL)
			counts++
		} else {
			amqpConn = connection

			break
		}

		if counts > _retryTimes {
			zap.S().Error("failed to retry", err)

			return nil, ErrCannotConnectRabbitMQ
		}

		zap.S().Info("Backing off for 2 seconds...")
		time.Sleep(_backOffSeconds * time.Second)

		continue
	}

	zap.S().Info("Connected to RabbitMQ 🎉")

	return amqpConn, nil
}
