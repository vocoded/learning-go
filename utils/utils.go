package utils

import (
  "fmt"
  "os"
  "bufio"
)

func GetInput(message string) string {
  fmt.Print(message)
  scanner := bufio.NewScanner(os.Stdin)
  if scanner.Scan() {
    return scanner.Text()
  }
  return ""
}
