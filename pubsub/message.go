package pubsub

type Message[T any] struct {
	Body T
}

func NewMessage[T any](msg T) *Message[T] {
	return &Message[T]{
		Body: msg,
	}
}
