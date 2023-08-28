package pubsub

import "sync"

type Subscribes[T any] map[string]*Subscribe[T]

type Broker[T any] struct {
	subscribers map[string]Subscribes[T]
	mutex       sync.Mutex
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		subscribers: make(map[string]Subscribes[T]),
	}
}

func (b *Broker[T]) Subscribe(topic string) *Subscribe[T] {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	sub := NewSubscribe[T](topic)

	if _, ok := b.subscribers[topic]; !ok {
		b.subscribers[topic] = make(Subscribes[T])
	}

	b.subscribers[topic][sub.id] = sub

	return sub
}

func (b *Broker[T]) Unsubscribe(sub *Subscribe[T]) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	delete(b.subscribers[sub.topic], sub.id)
}

func (b *Broker[T]) Publish(topic string, message T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, sub := range b.subscribers[topic] {
		sub.message <- message
	}
}
