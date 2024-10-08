package pokecache

import (
	"sync"
	"time"

)


type cacheEntry struct {
    createAt time.Time
    val  []byte
}
type Cache struct {
   data  map[string]cacheEntry
    mutex sync.RWMutex
    cleanupTick time.Duration
}


func NewCache(cleanupTick time.Duration) *Cache{
    cache := &Cache{
        data : make(map[string]cacheEntry),
        cleanupTick: cleanupTick,
        

    }
    go cache.reapLoop()
    return cache
}


func (c *Cache) Add(key string, val []byte){
    c.mutex.Lock()
    defer c.mutex.Unlock()
     c.data[key] = cacheEntry{
        val : val,
        createAt : time.Now().UTC(),
    }}

func (c *Cache) Get(key string) ([]byte,bool){
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if entry, ok := c.data[key]; ok {
      return entry.val,ok
    }
    return []byte{} ,false

}
func (c *Cache) reapLoop(){
    ticker := time.NewTicker(c.cleanupTick)
    defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.data {
			// Check if the entry is older than the cleanupTick interval
			if time.Since(entry.createAt) > c.cleanupTick {
				delete(c.data, key) // Remove stale entry
			}
		}
		c.mutex.Unlock()
}
}
