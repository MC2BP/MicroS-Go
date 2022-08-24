package cachelib

import (
	"sync"
	"time"
)

type CacheItem[K any] struct {
	lastUpdated time.Time
	value       K
}

type Cache[K comparable, L any] struct {
	sync.Mutex
	getter       func(K) *L
	values       map[K]CacheItem[L]
	refreshAfter time.Duration
}

func NewCache[K comparable, L any](fetch func(K) *L, refresh time.Duration) Cache[K, L] {
	return Cache[K, L]{
		getter:       fetch,
		refreshAfter: refresh,
	}
}

func (c *Cache[K, L]) Get(index K) *L {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// get value from cache
	v, has := c.values[index]
	if has && v.lastUpdated.Add(c.refreshAfter).After(time.Now()) {
		return &v.value
	}

	// value not in cache, or outdated
	value := c.getter(index)
	if value != nil {
		c.values[index] = CacheItem[L]{
			lastUpdated: time.Now(),
			value:       *value,
		}
	}

	return value
}
