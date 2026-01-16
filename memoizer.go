package cache

import (
	"errors"
	"sync"
	"time"
)

type Item struct {
	value     any
	ttl       time.Duration
	createdAt time.Time
}

type Cache struct {
	memo     map[string]Item
	mu       sync.RWMutex
	stopChan chan struct{}
}

func (c *Cache) checkExpiredValue() {
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	for {
		select {
		case <-c.stopChan:
			return
		case <-ticker.C:
			c.clear()
		}

	}
}

func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.memo {
		if time.Since(v.createdAt) > v.ttl {
			delete(c.memo, k)
		}
	}
}

func (c *Cache) Close() {
	close(c.stopChan)
}

func New() *Cache {
	c := &Cache{
		memo: make(map[string]Item),
	}
	go c.checkExpiredValue()
	return c
}

func (c *Cache) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.memo[key]
	if !ok {
		return Item{}, errors.New("cache item is not set")
	}
	if time.Since(item.createdAt) > item.ttl {
		delete(c.memo, key)
		return Item{}, errors.New("cache item is expired")
	}
	return item.value, nil
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.memo[key] = Item{value, ttl, time.Now()}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.memo, key)
}
