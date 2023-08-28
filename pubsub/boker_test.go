package pubsub

import "testing"

func TestSubscribe(t *testing.T) {
	broker := NewBroker[string]()

	sub := broker.Subscribe("topic")
	want := broker.subscribers["topic"][sub.id]

	if sub != want {
		t.Errorf("got %v, wanted %v", sub, want)
	}
}

func TestUnsubscribe(t *testing.T) {
	broker := NewBroker[string]()

	sub := broker.Subscribe("topic")
	broker.Unsubscribe(sub)

	if len(broker.subscribers["topic"]) != 0 {
		t.Errorf("got %v, wanted %v", len(broker.subscribers["topic"]), 0)
	}
}

func TestPublish(t *testing.T) {
	broker := NewBroker[string]()

	sub := broker.Subscribe("topic")

	var got string
	want := "message"

	sub.Receive(func(msg *Message[string]) {
		got = msg.Body
	})

	broker.Publish("topic", want)

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
