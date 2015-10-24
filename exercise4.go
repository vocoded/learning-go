package main

import (
  "fmt"
  "os" 
  "bufio"
  "strings"
)

type Matcher interface {
  Match(term string) []string
}

type FileMatcher struct {
  sourceFile string
}

func (m *FileMatcher) Match(term string) []string {
  snippets := make([]string, 0)
  
  f, err := os.Open(m.sourceFile)
  if err != nil {
    return snippets
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    m.MatchLine(scanner.Text(), term, &snippets)
  }
  
  return snippets  
}

func (m *FileMatcher) MatchLine(line string, term string, snippets *[]string) {
    previousTerm := ""
    for _, currentTerm := range strings.Split(line, " ") {
      if currentTerm == term {
        *snippets = append(*snippets, previousTerm + " " + currentTerm)
      }
      previousTerm = currentTerm      
    }  
}

func exercise4() {
  var term string
  matcher := FileMatcher {"terms.txt"}
  fmt.Print("Enter term to search: ")
  fmt.Scanln(&term)
  
  snippets := matcher.Match(term)
  fmt.Println("Found", len(snippets), "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}
