package pubsub

import (
	"testing"
)

func TestReceive(t *testing.T) {
	sub := NewSubscribe[string]()
	want := "some message"

	sub.Receive(func(msg *Message[string]) {
		if msg.Body != want {
			t.Errorf("got %s, wanted %s", msg.Body, want)
		}
	})

	sub.message <- want
}
