package main

import (
  "fmt"
  "os"
  "github.com/vocoded/learning-go/utils"
)

func WriteFile(strings []string) {
  file, err := os.Create("strings.txt")
  if err != nil {
    fmt.Println("Unable to create new file", err)
    return
  }
  defer file.Close()
  
  for _, str := range strings {
    file.WriteString(str)
    file.WriteString("\n")
  }
}

func FindUniqueSubstrings(input string, length int) []string {
  if length > len(input) {
    return []string {}
  }
  
  substrings := make(map[string]bool)
  
  for i := 0; i <= len(input) - length; i++ {
    str := input[i:i+length]
    substrings[str] = true
  }
  
  return utils.GetKeys(substrings)
}

// A little more string fun along with file creation
func exercise9() {
  input := utils.GetInput("Enter some text: ")
  substrings := FindUniqueSubstrings(input, 3)
  WriteFile(substrings)
}
