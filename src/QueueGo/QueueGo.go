package QueueGo

import "time"

type Message struct {
	Topic     string    `json:"topic"`
	Content   any       `json:"message"`
	Timestamp time.Time `json:"time"`
}
