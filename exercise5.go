package main

import (
  "fmt"
  "regexp"
  "strings"
)

type TermFinder interface {
  Find() []string
}

type FileTermFinder struct {
  sourceFile string
}

func (m *FileTermFinder) Find() []string {
  wordMap := make(map[string]bool)
  wordRegex, _ := regexp.Compile("\\W")

  findAction := func(term string) {
    // This is fairly rough; won't help us with hyphenated words or contractions
    term = wordRegex.ReplaceAllString(term, "")
    if term != "" {
      wordMap[strings.ToLower(term)] = true        
    }    
  }

  IterateFile(m.sourceFile, findAction)
  return Keys(wordMap)
}

func Keys(m map[string]bool) []string {
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

// Now we explore the map type a bit along with a gentle regular expression
func exercise5() {
  counter := FileTermFinder {"terms.txt"}  
  words := len(counter.Find())
  fmt.Println("Found", words, "unique words in file")
}
