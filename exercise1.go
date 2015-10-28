package main

import (
  "fmt"
)

// Simpler than FizzBuzz
func exercise1() {
  // Only one loop construct available in Go
  for i := 1; i <= 100; i++ {
    if i % 5 == 0 {
      fmt.Println(i)
    }
  }
}
