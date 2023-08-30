package QueueGo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func NewQueueGo() *QueueGo {
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

		subscribers, exists := b.subscribers[topic]
		if !exists {
			subscribers = []chan<- Message{}
			b.subscribers[topic] = subscribers
		}

		messages, exists := b.messages[topic]
		if !exists {
			messages = []Message{}
			b.messages[topic] = messages
		}

		b.mu.Unlock()

		b.mu.Lock()
		b.messages[topic] = append(messages, message)
		b.mu.Unlock()

		b.mu.Lock()
		for _, subscriber := range subscribers {
			subscriber <- message
		}
		b.mu.Unlock()

		err = b.saveToLog(topic, message)
		if err != nil {
			log.Println("Error saving to log:", err)
		}
	}
	return
}

func (b *QueueGo) saveToLog(topic string, message Message) error {
	fileName := fmt.Sprintf("%s.log", topic)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(message)
}

func (b *QueueGo) Subscribe(queue string) <-chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 100)

	b.subscribers[queue] = append(b.subscribers[queue], ch)

	for _, message := range b.messages[queue] {
		ch <- message
	}

	return ch
}

func main() {

	Queue := NewQueueGo()

	subscriber1 := Queue.Subscribe("topic1")
	subscriber2 := Queue.Subscribe("topic2")

	go func() {
		for message := range subscriber1 {
			log.Println("Subscriber 1 - Topic: ", message.Topic, " Content: ", message.Content, " Timestamp: ", message.Timestamp)
		}
	}()

	go func() {
		for message := range subscriber2 {
			log.Println("Subscriber 2 - Topic: ", message.Topic, " Content: ", message.Content, " Timestamp: ", message.Timestamp)
		}
	}()

	Queue.Publish("topic1", "Message 1 for Topic 1", true, []string{"message1", "message2"})
	Queue.Publish("topic2", "Message 1 for Topic 2", false, []int{1, 2, 3, 4, 5})

	time.Sleep(time.Second * 10)

}
