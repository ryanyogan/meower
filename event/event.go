package event

import "github.com/ryanyogan/meower/schema"

// Store defines the pub/sub events
type Store interface {
	Close()
	PublishMeowCreated(meow schema.Meow) error
	SubscribeMeowCreated() (<-chan MeowCreatedMessage, error)
	OnMeowCreated(f func(MeowCreatedMessage)) error
}

var impl Store

// SetEventStore -
func SetEventStore(es Store) {
	impl = es
}

// Close -
func Close() {
	impl.Close()
}

// PublishMeowCreated --
func PublishMeowCreated(meow schema.Meow) error {
	return impl.PublishMeowCreated(meow)
}

// SubscribeMeowCreated --
func SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	return impl.SubscribeMeowCreated()
}

// OnMeowCreated --
func OnMeowCreated(f func(MeowCreatedMessage)) error {
	return impl.OnMeowCreated(f)
}
