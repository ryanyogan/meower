package schema

import "time"

// Meow defines the payload with an ID, Body, and XreatedAt
type Meow struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
