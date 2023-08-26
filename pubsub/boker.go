package pubsub

import "sync"

type Broker[T any] struct {
	subscribers map[string][]*Subscribe[T]
	mutex       sync.Mutex
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		subscribers: make(map[string][]*Subscribe[T]),
	}
}

func (b *Broker[T]) Subscribe(topic string) *Subscribe[T] {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	sub := NewSubscribe[T]()
	b.subscribers[topic] = append(b.subscribers[topic], sub)

	return sub
}

func (b *Broker[T]) Publish(topic string, message T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, sub := range b.subscribers[topic] {
		sub.message <- message
	}
}
