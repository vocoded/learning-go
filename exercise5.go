package main

import (
  "fmt"
  "os"
  "bufio"
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

  f, err := os.Open(m.sourceFile)
  if err != nil {
    return 0
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    for _, currentTerm := range strings.Split(scanner.Text(), " ") {
      currentTerm = wordRegex.ReplaceAllString(currentTerm, "")
      if currentTerm != "" {
        wordMap[strings.ToLower(currentTerm)] = true        
      }
    }
  }

  return len(wordMap)
}

func exercise5() {
  counter := FileTermCounter {"terms.txt"}  
  words := counter.Count()
  fmt.Println("Found", words, "unique words in file")
}
