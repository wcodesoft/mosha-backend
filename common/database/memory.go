package database

import (
	"context"
	"fmt"
)

// MemoryDatabase is a database that stores data in memory.
type MemoryDatabase[T any] struct {
	// data is the data stored in the database.
	data map[string]*T
}

// NewMemoryDatabase creates a new MemoryDatabase.
func NewMemoryDatabase[T any]() MemoryDatabase[T] {
	return MemoryDatabase[T]{
		data: make(map[string]*T),
	}
}

// Create creates a new item in the database.
func (db MemoryDatabase[T]) Create(ctx context.Context, id string, item *T) (*T, error) {
	if _, ok := db.data[id]; ok {
		return nil, fmt.Errorf("item with id %s already exists", id)
	}
	db.data[id] = item
	return item, nil
}

// Read reads an item from the database.
func (db MemoryDatabase[T]) Read(ctx context.Context, id string) (*T, error) {
	item, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("item with id %s not found", id)
	}
	return item, nil
}

// Update updates an item in the database.
func (db MemoryDatabase[T]) Update(ctx context.Context, id string, item *T) (*T, error) {
	if _, ok := db.data[id]; !ok {
		return nil, fmt.Errorf("item with id %s not found", id)
	}
	db.data[id] = item
	return item, nil
}

// Delete deletes an item from the database.
func (db MemoryDatabase[T]) Delete(ctx context.Context, id string) error {
	if _, ok := db.data[id]; !ok {
		return fmt.Errorf("item with id %s not found", id)
	}
	delete(db.data, id)
	return nil
}

// List lists all items in the database.
func (db MemoryDatabase[T]) List(ctx context.Context) []*T {
	items := make([]*T, 0)
	for _, item := range db.data {
		items = append(items, item)
	}
	return items
}
