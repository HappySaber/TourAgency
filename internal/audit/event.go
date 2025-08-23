package audit

import (
	"encoding/json"
	"time"
)

type Event struct {
	EventID       string    `json:"event_id"`
	Event         string    `json:"event"`
	Entity        string    `json:"entity"`
	EntityID      string    `json:"entity_id"`
	ActorID       string    `json:"actor_id,omitempty"`
	CorrelationID string    `json:"correlation_id,omitempty"`
	IP            string    `json:"ip,omitempty"`
	UserAgent     string    `json:"user_agent,omitempty"`
	At            time.Time `json:"at"`
	Before        jsonRaw   `json:"before,omitempty"`
	After         jsonRaw   `json:"after,omitempty"`
}

type jsonRaw []byte

func MustMarshal(v interface{}) jsonRaw {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		// логировать можно здесь, если хочешь
		return nil
	}
	return jsonRaw(b)
}
