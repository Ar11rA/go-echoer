package utils

import (
	"context"
	"quote-server/constants"

	"github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient interface for better testability
type RabbitMQClient interface {
	Publish(ctx context.Context, exchange, key string, body []byte) error
}

// RabbitMQUtil is the implementation of RabbitMQClient
type RabbitMQUtil struct {
	Conn *amqp091.Connection
	Ch   *amqp091.Channel
}

// NewRabbitMQUtil creates a new RabbitMQUtil instance
func NewRabbitMQUtil(conn *amqp091.Connection) (*RabbitMQUtil, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_ = ch.ExchangeDeclare(
		constants.ExchangeName, // name
		constants.ExchangeType, // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	q, _ := ch.QueueDeclare(
		constants.QueueName, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	ch.QueueBind(
		q.Name,                 // queue name
		constants.RoutingKey,   // routing key
		constants.ExchangeName, // exchange name
		false,                  // no-wait
		nil,                    // arguments
	)
	return &RabbitMQUtil{Conn: conn, Ch: ch}, nil
}

// Publish publishes a message to the specified exchange and routing key
func (r *RabbitMQUtil) Publish(ctx context.Context, exchange, key string, body []byte) error {
	return r.Ch.Publish(exchange, key, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
