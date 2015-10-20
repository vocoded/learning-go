package main

import (
  "fmt"
  "os" 
  "bufio"
  "strings"
)

func GetSearchTerm() string {
  reader := bufio.NewReader(os.Stdin) 
  term, _ := reader.ReadString('\n')
  return term
}

func ScanSearchTerm() string {
  var term string
  fmt.Scanln(&term)
  return term
}

func MatchTerms(term string) (int, []string, error) {
  count := 0;
  snippets := make([]string, 0)
  
  f, err := os.Open("terms.txt")
  if err != nil {
    return 0, snippets, err
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    currentLine := scanner.Text()
    previousTerm := ""
    for _, currentTerm := range strings.Split(currentLine, " ") {
      if currentTerm == term {
        count = count + 1
        snippets = append(snippets, previousTerm + " " + currentTerm)
      }
      previousTerm = currentTerm      
    }
  }
  
  return count, snippets, nil
}

func exercise3() {
  fmt.Print("Enter term to search: ")
  term := ScanSearchTerm()
  count, snippets, _ := MatchTerms(term)
  fmt.Println("Found", count, "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}
