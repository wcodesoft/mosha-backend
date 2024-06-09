package database

import "context"

// Database is the interface that wraps the basic database operations.
type Database[T any] interface {
	// Create creates a new item in the database.
	Create(ctx context.Context, id string, item *T) (*T, error)
	// Read reads an item from the database.
	Read(ctx context.Context, id string) (*T, error)
	// Update updates an item in the database.
	Update(ctx context.Context, id string, item *T) (*T, error)
	// Delete deletes an item from the database.
	Delete(ctx context.Context, id string) error
}
