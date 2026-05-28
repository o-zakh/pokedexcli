package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	Mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// You'll probably want to expose a NewCache() function that creates a new cache
// with a configurable interval (time.Duration).
func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		Cache: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	cEntry, exists := c.Cache[key]
	if exists {
		return cEntry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Mu.Lock()
		for key, value := range c.Cache {
			if time.Since(value.createdAt) > interval {
				delete(c.Cache, key)
			}
		}
		c.Mu.Unlock()
	}
}
