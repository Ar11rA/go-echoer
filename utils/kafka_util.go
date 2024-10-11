package utils

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaClient interface {
	WriteMessages(ctx context.Context, message kafka.Message) error
}

type KafkaUtil struct {
	Writer *kafka.Writer
}

func NewKafkaUtil(Writer *kafka.Writer) *KafkaUtil {
	return &KafkaUtil{Writer: Writer}
}

func (k *KafkaUtil) WriteMessages(ctx context.Context, message kafka.Message) error {
	err := k.Writer.WriteMessages(ctx, message)
	return err
}
