package main

import (
	"fmt"
	"time"

	"github.com/andreglatz/pubsub/pubsub"
)

func main() {
	broker := pubsub.NewBroker[string]()

	sub1 := broker.Subscribe("topic1")
	sub2 := broker.Subscribe("topic1")

	sub1.Receive(func(msg *pubsub.Message[string]) {
		fmt.Println("sub1", msg.Body)
	})

	sub2.Receive(func(msg *pubsub.Message[string]) {
		fmt.Println("sub2", msg.Body)
	})

	broker.Publish("topic1", "{\"message\": \"hello world!\"}")
	broker.Publish("topic1", "{\"message\": \"hello world!\"}")

	broker.Unsubscribe(sub1)

	broker.Publish("topic1", "{\"message\": \"hello world!\"}")
	broker.Publish("topic1", "{\"message\": \"hello world!\"}")

	time.Sleep(1 * time.Second)
}
