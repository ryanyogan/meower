package event

import (
	"bytes"
	"encoding/gob"

	"github.com/nats-io/nats.go"
	"github.com/ryanyogan/meower/schema"
)

// NatsEventStore --
type NatsEventStore struct {
	nc                      *nats.Conn
	meowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
}

// NewNats --
func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

// Close --
func (n *NatsEventStore) Close() {
	if n.nc != nil {
		n.nc.Close()
	}
	if n.meowCreatedSubscription != nil {
		n.meowCreatedSubscription.Unsubscribe()
	}
	close(n.meowCreatedChan)
}

// PublishMeowCreated --
func (n *NatsEventStore) PublishMeowCreated(meow schema.Meow) error {
	m := MeowCreatedMessage{meow.ID, meow.Body, meow.CreatedAt}
	data, err := n.writeMessage(&m)
	if err != nil {
		return err
	}
	return n.nc.Publish(m.Key(), data)
}

func (*NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// OnMeowCreated --
func (n *NatsEventStore) OnMeowCreated(f func(MeowCreatedMessage)) (err error) {
	m := MeowCreatedMessage{}
	n.meowCreatedSubscription, err = n.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		n.readMessage(msg.Data, &m)
		f(m)
	})

	return
}

func (*NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
