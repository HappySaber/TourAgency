package audit

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func HandleAuditEvent(ctx context.Context, msg kafka.Message, repo Repository) {
	var event Event
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		log.Printf("failed to unmarshal audit event: %v", err)
		return
	}

	if err := repo.SaveEvent(ctx, event); err != nil {
		log.Printf("failed to save audit event: %v", err)
	}
}
