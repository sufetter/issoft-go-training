package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	capacity int
	items    map[string]*cacheItem
	keys     []string
	mutex    sync.RWMutex
}

type cacheItem struct {
	data         any
	lastAccessed time.Time
}

func NewCache(capacity int) (*Cache, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity provided")
	}

	c := &Cache{
		capacity: capacity,
		items:    make(map[string]*cacheItem),
		keys:     make([]string, 0, capacity),
	}

	go c.startEvictionWorker()
	return c, nil
}

func (c *Cache) startEvictionWorker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.evictExpiredItems()
	}
}

func (c *Cache) evictExpiredItems() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	currentTime := time.Now()
	for key, item := range c.items {
		if currentTime.Sub(item.lastAccessed) > 10*time.Second {
			delete(c.items, key)
			c.removeKey(key)
		}
	}
}

func (c *Cache) Set(key string, data any) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(key) == 0 {
		return errors.New("empty key provided")
	}

	if item, exists := c.items[key]; exists {
		item.data = data
		item.lastAccessed = time.Now()
		return nil
	} else {
		if len(c.items) >= c.capacity {
			c.evictOldestItem()
		}
		c.items[key] = &cacheItem{data: data, lastAccessed: time.Now()}
		c.keys = append(c.keys, key)
	}

	return nil
}

func (c *Cache) Get(key string) (any, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if len(key) == 0 {
		return nil, errors.New("empty key provided")
	}

	if item, exists := c.items[key]; exists {
		item.lastAccessed = time.Now()
		return item.data, nil
	}
	return nil, errors.New("key not found")
}

func (c *Cache) Remove(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(key) == 0 {
		return errors.New("empty key provided")
	}

	if _, exists := c.items[key]; exists {
		delete(c.items, key)
		c.removeKey(key)
		return nil
	}
	return errors.New("key not found")
}

func (c *Cache) evictOldestItem() {
	//Функционал схож с методом выше, но
	//я посчитал, что все же это довольно
	//разные методы и решил их не объединять в 1

	if len(c.items) == 0 || len(c.keys) == 0 {
		return
	}

	oldestKey := c.keys[0]
	oldestAccessed := c.items[oldestKey].lastAccessed

	for _, key := range c.keys {
		if c.items[key].lastAccessed.Before(oldestAccessed) {
			oldestKey = key
			oldestAccessed = c.items[key].lastAccessed
		}
	}

	delete(c.items, oldestKey)
	c.removeKey(oldestKey)
}

func (c *Cache) removeKey(key string) {
	for i, k := range c.keys {
		if k == key {
			c.keys = append(c.keys[:i], c.keys[i+1:]...)
			break
		}
	}
}
