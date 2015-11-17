package main

import (
  "fmt"
  "github.com/vocoded/learning-go/utils"
)

func Reverse(input string) string {
  length := len(input)
  reverse := make([]rune, length, length)
  for i, c := range input {
    reverse[length - i - 1] = c
  }
  
  return string(reverse)
}

// The old standby - reverse a string
func exercise8() {
  input := utils.GetDefaultInput()
  reverse := Reverse(input)
  fmt.Println("The reverse is", reverse)
}
