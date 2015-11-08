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

func GetKeys(m map[string]bool) []string {
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}
