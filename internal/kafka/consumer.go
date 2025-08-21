package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

func (kc *KafkaConsumer) Start(ctx context.Context, handler func(msg kafka.Message)) {
	go func() {
		defer kc.reader.Close()
		for {
			m, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Kafka consumer error: %v", err)
				break
			}
			handler(m)
		}
	}()
}
