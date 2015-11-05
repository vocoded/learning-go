package main

import (
  "fmt"
  "strings"
)

type Matcher interface {
  Match(term string) []string
}

type FileMatcher struct {
  sourceFile string
}

func (m *FileMatcher) Match(term string) []string {
  return MatchTermsFromFile(m.sourceFile, term)
}

func MatchTermsViaClass(term string) {
  matcher := FileMatcher {"terms.txt"}
  snippets := matcher.Match(strings.ToLower(term))
  fmt.Println("Found", len(snippets), "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}

// Building on the previous application, this adds a class-based implementation
func exercise4() {
  var term string
  fmt.Print("Enter term to search: ")
  if _, err := fmt.Scanln(&term); err == nil {
    MatchTermsViaClass(term)
  }  
}
