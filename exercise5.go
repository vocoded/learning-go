package main

import (
  "fmt"
  "regexp"
  "strings"
)

type TermCounter interface {
  Count() int
}

type FileTermCounter struct {
  sourceFile string
}

func (m *FileTermCounter) Count() int {
  wordMap := make(map[string]bool)
  wordRegex, _ := regexp.Compile("\\W")

  countAction := func(term string) {
    term = wordRegex.ReplaceAllString(term, "")
    if term != "" {
      wordMap[strings.ToLower(term)] = true        
    }    
  }

  IterateFile(m.sourceFile, countAction)
  return len(wordMap)
}

func exercise5() {
  counter := FileTermCounter {"terms.txt"}  
  words := counter.Count()
  fmt.Println("Found", words, "unique words in file")
}
