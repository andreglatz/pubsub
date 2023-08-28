package pubsub

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Subscribes[T any] map[string]*Subscribe[T]

type Broker[T any] struct {
	subscribers map[string]Subscribes[T]
	mutex       sync.Mutex
	path        string
}

func NewBroker[T any](path string) *Broker[T] {
	return &Broker[T]{
		subscribers: make(map[string]Subscribes[T]),
		path:        path,
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
		go b.writeToFile(topic, message)
		sub.message <- message
	}
}

func (b *Broker[T]) writeToFile(topic string, message T) {
	date := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s/%s-%s.log", b.path, topic, date)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dateTime := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s %v\n", dateTime, message)

	if _, err := f.WriteString(log); err != nil {
		panic(err)
	}
}
