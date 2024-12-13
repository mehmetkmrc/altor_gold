package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Option func(*Consumer)

func QueueName(queueName string) Option {
	return func(p *Consumer) {
		p.QueueName = queueName
	}
}

func QueueArg(queueArg amqp.Table) Option {
	return func(p *Consumer) {
		p.QueueArg = queueArg
	}
}

func ConsumerTag(consumerTag string) Option {
	return func(p *Consumer) {
		p.ConsumerTag = consumerTag
	}
}

func WorkerPoolSize(workerPoolSize int) Option {
	return func(p *Consumer) {
		p.WorkerPoolSize = workerPoolSize
	}
}
