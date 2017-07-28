package cache

import (
	"sync"
	"time"

	"github.com/blainsmith/figment/backend"
)

// Cache is the abstract layer around the Backend cache interface. It keeps of Items and their lifecycle functions.
type Cache struct {
	mu      sync.RWMutex
	backend backend.Backend
	items   map[string]Item
}

// Item holds the lifecycle information of individual cache items.
type Item struct {
	Key       string
	trigger   []ItemFunc
	before    []ItemFunc
	after     []ItemFunc
	CreatedAt time.Time
}

// New creates a new Cache with a Backend cache implementation.
func New(b backend.Backend) *Cache {
	c := Cache{
		backend: b,
		items:   make(map[string]Item),
	}
	return &c
}

// Set adds an item to the cache with optional lifecycle Options
func (c *Cache) Set(key string, value []byte, options ...Option) error {
	c.backend.Set(key, value)

	item := Item{
		Key:       key,
		CreatedAt: time.Now(),
	}
	for _, option := range options {
		option(c, &item)
	}

	c.mu.Lock()
	c.items[key] = item
	c.mu.Unlock()

	return nil
}

// Get pulls an item from the cache and executes all lifecycle options set on the item.
func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	item := c.items[key]
	c.mu.RUnlock()

	for _, f := range item.trigger {
		go func(f ItemFunc) {
			f(&item)
		}(f)
	}
	for _, f := range item.before {
		f(&item)
	}
	backendItem, err := c.backend.Get(key)
	if err != nil {
		return nil, err
	}
	for _, f := range item.after {
		f(&item)
	}
	return backendItem.([]byte), nil
}

// Delete removes an item from the cache.
func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()

	c.backend.Delete(key)

	return nil
}
