package main

import (
  "fmt"
  "math/rand"
  "time"
)

func FindRandomMultiple(c chan int) {
  rng := rand.New(rand.NewSource(time.Now().UnixNano())) 
  for {
    num := rng.Intn(1000)
    if num % 99 == 0 {
      c <- num
      break
    }
  }  
}

// Random values and basic concurrency
func exercise10() {
  channel := make(chan int)
  
  go FindRandomMultiple(channel)
  
  result := <-channel
  fmt.Println(result)
}
