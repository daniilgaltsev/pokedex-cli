package cache

import (
	"sync"
	"time"
)

type CacheEntry struct{
	value []byte
	createdAt time.Time
}

type Cache struct{
	data map[string]CacheEntry
	retention time.Duration
	mutex sync.RWMutex
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	c.data[key] = CacheEntry{
		value: value,
		createdAt: time.Now(),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	entry, ok := c.data[key]
	c.mutex.RUnlock()
	return entry.value, ok
}

func (c *Cache) CheckRetention() {
	c.mutex.Lock()
	for key, entry := range c.data {
		if time.Since(entry.createdAt) > c.retention {
			delete(c.data, key)
		}
	}
	c.mutex.Unlock()
}


func NewCache(retention time.Duration) Cache {
	cache := Cache{
		data: make(map[string]CacheEntry),
		retention: retention,
	}

	go func() { // NOTE: should ideally be disableable and tracked directly in the cache struct
		timer := time.NewTicker(retention)
		for {
			<-timer.C
			cache.CheckRetention()
		}
	}()

	return cache
}
