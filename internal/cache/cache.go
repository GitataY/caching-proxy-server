package cache

import (
    "sync"
    "time"
)

type CacheItem struct {
    Data       []byte   
    Expiration time.Time
}

type Cache struct {
    items map[string]CacheItem
    mu    sync.RWMutex 
    ttl   time.Duration 
}

func NewCache(ttl time.Duration) *Cache {
    cache := &Cache{
        items: make(map[string]CacheItem),
        ttl:   ttl,
    }
    
    go cache.cleanup()
    return cache
}

func (c *Cache) Get(url string) ([]byte, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, found := c.items[url]
    if !found {
        return nil, false
    }
   
    if time.Now().After(item.Expiration) {
        return nil, false
    }
    return item.Data, true
}

func (c *Cache) Set(url string, data []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[url] = CacheItem{
        Data:       data,
        Expiration: time.Now().Add(c.ttl),
    }
}

func (c *Cache) Clear() {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items = make(map[string]CacheItem) 
}

func (c *Cache) cleanup() {
    for {
        time.Sleep(time.Minute)
        c.mu.Lock()
        now := time.Now()
        for url, item := range c.items {
            if now.After(item.Expiration) {
                delete(c.items, url)
            }
        }
        c.mu.Unlock()
    }
}
