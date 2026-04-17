package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.Mutex
	cache map[string]Item
}

type Item struct {
	value any
	ttl   time.Time
}

func New() *Cache {
	return &Cache{
		new(sync.Mutex),
		make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = Item{value, time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.cache[key]

	if !ok {
		return nil, errors.New("key not found")
	}

	if time.Now().After(item.ttl) {
		c.delete(key)
		return nil, errors.New("key is expired")
	}

	return item.value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.cache, key)
}

func (c *Cache) delete(key string) {
	delete(c.cache, key)
}
