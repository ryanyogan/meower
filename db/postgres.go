package db

import (
	"context"
	"database/sql"

	"github.com/ryanyogan/meower/schema"
)

// PostgresRepository holds a connection to a live postgres conn
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgres creates a new connection to a PG DB based on the host url
// that is provided.
func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

// Close wil close the db connection
func (r *PostgresRepository) Close() {
	r.db.Close()
}

// InsertMeow will insert a new meow into the db or return an error
func (r *PostgresRepository) InsertMeow(ctx context.Context, meow schema.Meow) error {
	_, err := r.db.Exec(
		"INSERT INTO meows(id, body, created_at) VALUES($1, $2, $3)",
		meow.ID, meow.Body, meow.CreatedAt,
	)

	return err
}

// ListMeows will query the db and return an array of meows, or an error
func (r *PostgresRepository) ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	rows, err := r.db.Query("SELECT * FROM meows ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meows := []schema.Meow{}
	for rows.Next() {
		meow := schema.Meow{}
		if err := rows.Scan(&meow.ID, &meow.Body, &meow.CreatedAt); err == nil {
			meows = append(meows, meow)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return meows, nil
}
