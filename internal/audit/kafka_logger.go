package audit

import (
	"TurAgency/internal/kafka"
	"context"
	"encoding/json"
	"fmt"
)

type Logger interface {
	Log(ctx context.Context, evt Event)
}

type kafkaLogger struct {
	producer *kafka.KafkaProducer
}

func NewKafkaLogger(p *kafka.KafkaProducer) *kafkaLogger {
	return &kafkaLogger{producer: p}
}

func (kl *kafkaLogger) Log(ctx context.Context, evt Event) error {
	b, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	key := []byte(fmt.Sprintf("%s:%s", evt.Entity, evt.EventID))
	return kl.producer.SendMessage(ctx, key, b)
}
