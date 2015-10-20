package main

import (
  "fmt"
  "strings"
  "strconv"
)

func FindMultiples(limit int, divisor int) string {
  multiples := []string {}
  for i := 1; i <= limit; i++ {
    if i % divisor == 0 {
      multiples = append(multiples, strconv.Itoa(i))
    }
  }
  // My kingdom for a map function!
  return strings.Join(multiples, ", ")
}

func exercise2() {
  limit, divisor := 200, 10
  fmt.Println("Multiples of", divisor, "up to", limit, ":", FindMultiples(limit, divisor))
}
