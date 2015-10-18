package main

import (
  "fmt"
)

func exercise1() {
  for i := 1; i <= 100; i++ {
    if i % 5 == 0 {
      fmt.Println(i)
    }
  }
}
