package services

import (
	"context"
	"quote-server/utils"

	"github.com/segmentio/kafka-go"
)

// KafkaService interface for Kafka-related operations
type KafkaService interface {
	PublishMessage(ctx context.Context, key string, message string) error
}

// KafkaServiceImpl implements KafkaService
type KafkaServiceImpl struct {
	KafkaUtil utils.KafkaClient
}

// NewKafkaService creates a new KafkaServiceImpl
func NewKafkaService(kafkaUtil utils.KafkaClient) KafkaService {
	return &KafkaServiceImpl{
		KafkaUtil: kafkaUtil,
	}
}

// PublishMessage sends a message to Kafka
func (k *KafkaServiceImpl) PublishMessage(ctx context.Context, key string, message string) error {
	// Prepare the Kafka message with key and value
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	}

	// Publish the message using the Kafka utility
	return k.KafkaUtil.WriteMessages(ctx, kafkaMessage)
}
