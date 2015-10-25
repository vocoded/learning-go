package main

import (
  "fmt"
  "os" 
  "bufio"
  "strings"
)

func GetSearchTerm() string {
  var term string
  fmt.Scanln(&term)
  return term
}

func MatchLine(line string, term string, snippets *[]string) {
    previousTerm := ""
    for _, currentTerm := range strings.Split(line, " ") {
      if strings.ToLower(currentTerm) == term {
        *snippets = append(*snippets, previousTerm + " " + currentTerm)
      }
      previousTerm = currentTerm      
    }  
}

func MatchTermsFromFile(file string, term string) []string {
  snippets := make([]string, 0)
  term = strings.ToLower(term)
  
  f, err := os.Open(file)
  if err != nil {
    return snippets
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    MatchLine(scanner.Text(), term, &snippets)
  }
  
  return snippets
}

func exercise3() {
  fmt.Print("Enter term to search: ")
  term := GetSearchTerm()
  snippets := MatchTermsFromFile("terms.txt", term)
  fmt.Println("Found", len(snippets), "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}
