package pubsub

import "github.com/andreglatz/pubsub/uuid"

type Subscribe[T any] struct {
	id      string
	topic   string
	message chan T
}

func NewSubscribe[T any](topic string) *Subscribe[T] {
	return &Subscribe[T]{
		id:      uuid.NewV4(),
		topic:   topic,
		message: make(chan T),
	}
}

func (s *Subscribe[T]) Receive(callback func(msg *Message[T])) {
	go func() {
		for msg := range s.message {
			callback(NewMessage(msg))
		}
	}()
}
