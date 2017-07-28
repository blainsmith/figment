package backend

import (
	"errors"
	"sync"
)

type mapBackend struct {
	mu    sync.RWMutex
	items map[interface{}]interface{}
}

// Map creates a new Backend implementation using a map and a mutex.
func Map() Backend {
	return &mapBackend{
		items: make(map[interface{}]interface{}),
	}
}

// Get pulls and item from the map or returns an error if it does not exist.
func (m *mapBackend) Get(k interface{}) (interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if item, ok := m.items[k]; ok {
		return item, nil
	}
	return nil, errors.New("not found")
}

// Set adds an item to the map.
func (m *mapBackend) Set(k, v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[k] = v
	return nil
}

// Delete removes an item from the map.
func (m *mapBackend) Delete(k interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, k)
	return nil
}
