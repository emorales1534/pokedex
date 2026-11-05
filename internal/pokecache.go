package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	caches map[string]cacheEntry
	sync.Mutex 
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{caches: make(map[string]cacheEntry), interval: interval}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add (key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	c.caches[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	val, ok := c.caches[key]
	if !ok {
		return nil, false
	}
	return val, true
}

func (c *Cache) reapLoop () {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.Lock()
		for key, val := range c.caches {
			if time.Since(val.createdAt) > c.interval {
				delete(c.caches, key)
			}
		}
		c.Unlock()
	}
}
