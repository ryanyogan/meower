package db

import "context"

import "github.com/ryanyogan/meower/schema"

// Repository defines what you may do with the Data Layer
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

// SetRepository allows for dynamic injection of db
func SetRepository(repository Repository) {
	impl = repository
}

// Close the db connection
func Close() {
	impl.Close()
}

// InsertMeow inserts one meow into the db
func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

// ListMeows returns all meows from the db
func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
