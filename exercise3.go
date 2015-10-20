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

func MatchLine(line string, term string, snippets *[]string) {
    previousTerm := ""
    for _, currentTerm := range strings.Split(line, " ") {
      if currentTerm == term {
        *snippets = append(*snippets, previousTerm + " " + currentTerm)
      }
      previousTerm = currentTerm      
    }  
}

func MatchTerms(term string) []string {
  snippets := make([]string, 0)
  
  f, err := os.Open("terms.txt")
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
  term := ScanSearchTerm()
  snippets := MatchTerms(term)
  fmt.Println("Found", len(snippets), "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}
