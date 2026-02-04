package main

import (
  "bytes"
  "fmt"
  "sync"
)

func main() {
  // Create a pool of byte buffers
  bufferPool := sync.Pool{
    New: func() interface{} {
      fmt.Println("Creating new buffer")
      return new(bytes.Buffer)
    },
  }

  // Get buffer from pool
  buf1 := bufferPool.Get().(*bytes.Buffer)
  buf1.WriteString("Hello, ")
  buf1.WriteString("World!")
  fmt.Println("Buffer 1:", buf1.String())

  // Reset and return to pool
  buf1.Reset()
  bufferPool.Put(buf1)

  // Get another buffer (reuses the one we put back)
  buf2 := bufferPool.Get().(*bytes.Buffer)
  buf2.WriteString("Reused buffer!")
  fmt.Println("Buffer 2:", buf2.String())

  // Demonstrate with goroutines
  var wg sync.WaitGroup
  for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      buf := bufferPool.Get().(*bytes.Buffer)
      buf.Reset()
      buf.WriteString(fmt.Sprintf("Goroutine %d", id))
      fmt.Println(buf.String())
      bufferPool.Put(buf)
    }(i)
  }

  wg.Wait()
  fmt.Println("Done!")
}

