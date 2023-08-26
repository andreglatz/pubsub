package pubsub

type Subscribe[T any] struct {
	message chan T
}

func NewSubscribe[T any]() *Subscribe[T] {
	return &Subscribe[T]{
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
