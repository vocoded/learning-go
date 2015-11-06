package main

import (
  "fmt"
  "github.com/vocoded/learning-go/utils"
)

// The old standby - reverse a string
func exercise8() {
  input := utils.GetInput("Enter some text: ")
  length := len(input)
  reverse := make([]rune, length, length)
  for i, c := range input {
    reverse[length - i - 1] = c
  }
  fmt.Println("The reverse is", string(reverse))
}
