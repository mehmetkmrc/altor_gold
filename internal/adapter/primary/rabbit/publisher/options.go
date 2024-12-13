package publisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Option func(*Publisher)

func SetExchangeName(exchangeName string) Option {
	return func(p *Publisher) {
		p.ExchangeName = exchangeName
	}
}

func SetBindingKey(bindingKey string) Option {
	return func(p *Publisher) {
		p.BindingKey = bindingKey
	}
}

func SetMessageTypeName(messageTypeName string) Option {
	return func(p *Publisher) {
		p.MessageTypeName = messageTypeName
	}
}

func QueueArg(queueArg amqp.Table) Option {
	return func(p *Publisher) {
		p.QueueArg = queueArg
	}
}

func SetQueueNames(queues QueueNames) Option {
	return func(p *Publisher) {
		p.Queues = queues
	}
}
