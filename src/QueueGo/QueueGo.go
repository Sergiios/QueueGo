package QueueGo

import (
	"sync"
	"time"
)

type Message struct {
	Topic     string    `json:"topic"`
	Content   any       `json:"message"`
	Timestamp time.Time `json:"time"`
}

type QueueGo struct {
	subscribers map[string][]chan<- Message
	messages    map[string][]Message
	mu          sync.RWMutex
}

func New() *QueueGo {
	return &QueueGo{
		subscribers: make(map[string][]chan<- Message),
		messages:    make(map[string][]Message),
		mu:          sync.RWMutex{},
	}
}

func (b *QueueGo) Publish(topic string, contents ...any) (err error) {
	for _, content := range contents {
		message := Message{
			Topic:     topic,
			Content:   content,
			Timestamp: time.Now(),
		}

		b.mu.Lock()
		defer b.mu.Unlock()

		if _, exists := b.subscribers[topic]; !exists {
			b.subscribers[topic] = []chan<- Message{}
		}

		b.messages[topic] = append(b.messages[topic], message)

		for _, subscriber := range b.subscribers[topic] {
			subscriber <- message
		}
	}
	return
}

func (b *QueueGo) Subscribe(queue string) <-chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 100) // Buffered channel to avoid blocking

	b.subscribers[queue] = append(b.subscribers[queue], ch)

	// Send existing messages to the subscriber
	for _, message := range b.messages[queue] {
		ch <- message
	}

	return ch
}
