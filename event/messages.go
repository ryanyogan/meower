package event

import "time"

// Message defines that all messages must have a specific key to listen to
// returned from a Key() function e.g. "meow.created"
type Message interface {
	Key() string
}

// MeowCreatedMessage is the message payload upon the creation of a meow
type MeowCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

// Key defines the key for messages to listen to
func (m *MeowCreatedMessage) Key() string {
	return "meow.created"
}
