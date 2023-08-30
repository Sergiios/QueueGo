<div style="text-align: center;">
  <h1>GoQueue 📨</h1>
</div>

A Biblioteca QueueGo simula um sistema de fila em Go. A estrutura possibilita a publicação e inscrição em tópicos, enquanto também armazena as mensagens em arquivos de log separados por tópico.

<div style="text-align: center;">
  <h2>Como utilizar ?</h2>
</div>

Pré requisitos:

[Go](https://go.dev/)

Para adicionar a biblioteca ao seu projeto, basta fazer o seguinte comando:

```shell 
go get "github.com/Sergiios/QueueGo"
```

Exemplo de uso:

```Go
package main

import (
	"fmt"
	"time"

	QueueGo "github.com/Sergiios/QueueGo"
)

func main() {

	Queue := QueueGo.NewQueueGo()

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
