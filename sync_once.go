package main

import (
  "fmt"
  "sync"
)

var (
  instance *Database
  once     sync.Once
)

type Database struct {
  connection string
}

func GetDatabase() *Database {
  once.Do(func() {
    fmt.Println("Initializing database connection...")
    instance = &Database{connection: "postgres://localhost:5432"}
  })
  return instance
}

func main() {
  var wg sync.WaitGroup

  // Multiple goroutines trying to get database
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      db := GetDatabase()
      fmt.Printf("Goroutine %d got: %s\n", id, db.connection)
    }(i)
  }

  wg.Wait()

  // Verify same instance
  db1 := GetDatabase()
  db2 := GetDatabase()
  fmt.Printf("\nSame instance? %v\n", db1 == db2)
}

