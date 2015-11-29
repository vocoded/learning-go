package main

import (
  "fmt"
  "regexp"
  "strings"
)

var wordRegex, _ = regexp.Compile("\\W")

type TermFinder interface {
  Find() []string
}

type FileTermFinder struct {
  sourceFile string
  wordMap map[string]bool
}

func (m FileTermFinder) MatchTerm(term string) {
  // This is fairly rough; won't help us with hyphenated words or contractions
  term = wordRegex.ReplaceAllString(term, "")
  if term != "" {
    m.wordMap[strings.ToLower(term)] = true        
  }  
}

func (m FileTermFinder) Find() []string {
  if err := IterateFile(m.sourceFile, m.MatchTerm); err != nil {
    return []string {}
  }

  return Keys(m.wordMap)
}

// Oh my is this painful - sadly Go wasn't built to be expressive
func Keys(m map[string]bool) []string {
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

func FindFileTerms(finder TermFinder) {
  words := len(finder.Find())
  fmt.Println("Found", words, "unique words in file")
}

// Explore the map type a bit along with a gentle regular expression
func exercise5() {
  finder := FileTermFinder {"terms.txt", make(map[string]bool)}
  FindFileTerms(finder)
}
