package search

import (
	"context"

	"github.com/ryanyogan/meower/schema"
)

// Repository defines search actions
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ScanMeows(ctx context.Context, query string, skip, take uint64) ([]schema.Meow, error)
}

var impl Repository

// SetRepository --
func SetRepository(repo Repository) {
	impl = repo
}

// Close --
func Close() {
	impl.Close()
}

// InsertMeow --
func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

// ScanMeows --
func ScanMeows(ctx context.Context, query string, skip, take uint64) ([]schema.Meow, error) {
	return impl.ScanMeows(ctx, query, skip, take)
}
