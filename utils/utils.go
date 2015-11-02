package utils

import "fmt"

func GetInput(message string) string {
  var input string
  fmt.Print(message)
  fmt.Scanln(&input)
  return input
}
