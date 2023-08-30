<div style="text-align: center;">
  <h1>GoQueue 📨</h1>
</div>


O objetivo principal é criar uma biblioteca simples de Publish/Subscribe em Go.

<div style="text-align: center;">
  <h2>Como utilizar ? 😃</h2>
</div>

Pré requisitos:

[Go](https://go.dev/)

Para adicionar a biblioteca ao seu projeto, basta fazer o seguinte comando:

```shell 
go get "github.com/Sergiios/QueueGo"
```

Exemplo de uso:

```Go  
fuc main() {

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
  ```