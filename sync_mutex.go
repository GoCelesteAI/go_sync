package main

import (
  "fmt"
  "sync"
)

// Counter with mutex protection
type Counter struct {
  mu    sync.Mutex
  value int
}

func (c *Counter) Increment() {
  c.mu.Lock()
  c.value++
  c.mu.Unlock()
}

func (c *Counter) Value() int {
  c.mu.Lock()
  defer c.mu.Unlock()
  return c.value
}

func main() {
  counter := &Counter{}
  var wg sync.WaitGroup

  // Launch 100 goroutines
  for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for j := 0; j < 1000; j++ {
        counter.Increment()
      }
    }()
  }

  wg.Wait()
  fmt.Println("Final count:", counter.Value())
  fmt.Println("Expected:", 100*1000)
}

