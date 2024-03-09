package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    *sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}

} 

func (c *Cache) Get(key string)([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) ReadLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)){
			delete(c.entries, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go cache.ReadLoop(interval)
	return cache
}
