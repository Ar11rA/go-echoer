package services

import (
	"context"
	"quote-server/constants"
	"quote-server/utils"
)

// RabbitMQService interface for the RabbitMQ service
type RabbitMQService interface {
	SendMessage(ctx context.Context, body []byte) error
}

// RabbitMQServiceImpl implements RabbitMQService
type RabbitMQServiceImpl struct {
	RabbitMQUtil utils.RabbitMQClient
}

// NewRabbitMQService creates a new RabbitMQServiceImpl
func NewRabbitMQService(rabbitMQUtil utils.RabbitMQClient) RabbitMQService {
	return &RabbitMQServiceImpl{
		RabbitMQUtil: rabbitMQUtil,
	}
}

// SendMessage sends a message to RabbitMQ
func (r *RabbitMQServiceImpl) SendMessage(ctx context.Context, body []byte) error {
	return r.RabbitMQUtil.Publish(
		ctx,
		constants.ExchangeName,
		constants.RoutingKey,
		body,
	)
}
