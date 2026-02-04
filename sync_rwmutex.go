package main

import (
  "fmt"
  "sync"
  "time"
)

// Cache with read-write mutex
type Cache struct {
  mu   sync.RWMutex
  data map[string]string
}

func NewCache() *Cache {
  return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
  c.mu.RLock()
  defer c.mu.RUnlock()
  val, ok := c.data[key]
  return val, ok
}

func (c *Cache) Set(key, value string) {
  c.mu.Lock()
  defer c.mu.Unlock()
  c.data[key] = value
}

func main() {
  cache := NewCache()
  var wg sync.WaitGroup

  // Writer goroutine
  wg.Add(1)
  go func() {
    defer wg.Done()
    for i := 0; i < 5; i++ {
      key := fmt.Sprintf("key%d", i)
      cache.Set(key, fmt.Sprintf("value%d", i))
      fmt.Printf("Write: %s\n", key)
      time.Sleep(100 * time.Millisecond)
    }
  }()

  // Multiple reader goroutines
  for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      for j := 0; j < 5; j++ {
        key := fmt.Sprintf("key%d", j)
        if val, ok := cache.Get(key); ok {
          fmt.Printf("Reader %d: %s=%s\n", id, key, val)
        }
        time.Sleep(50 * time.Millisecond)
      }
    }(i)
  }

  wg.Wait()
  fmt.Println("Done!")
}

