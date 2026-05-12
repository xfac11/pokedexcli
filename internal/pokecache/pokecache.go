package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time //When the entry was created
	val       []byte    //Raw data
}
type Cache struct {
	storage map[string]cacheEntry
	mu      sync.Mutex
}

func (cache *Cache) Add(key string, val []byte) error {
	if key == "" {
		return fmt.Errorf("Cannot use an empty string as a key")
	}
	if val == nil {
		return fmt.Errorf("Val has to be a valid slice, not nil")
	}
	if len(val) == 0 {
		return fmt.Errorf("val has to contain something, not zero length")
	}
	cache.mu.Lock()
	cache.storage[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mu.Unlock()
	return nil
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.storage[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		for key, val := range cache.storage {
			elapsed := time.Since(val.createdAt)
			if elapsed > interval {
				cache.mu.Lock()
				delete(cache.storage, key)
				cache.mu.Unlock()
			}
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		storage: map[string]cacheEntry{},
		mu:      sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return &cache
}
