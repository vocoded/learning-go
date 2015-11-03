package main

import (
  "fmt"
  "os" 
  "bufio"
  "strings"
)

func GetSearchTerm() (string, error) {
  var term string
  if _, err := fmt.Scanln(&term); err != nil {
    fmt.Println("Please enter a single term")
    return "", err
  }
  return term, nil
}

func MatchTermsFromFile(file string, term string) []string {
  snippets := make([]string, 0)
  term = strings.ToLower(term)

  previousTerm := ""  
  matchAction := func(currentTerm string) {
    if strings.ToLower(currentTerm) == term {
      snippets = append(snippets, previousTerm + " " + currentTerm)
    }
    previousTerm = currentTerm    
  }

  IterateFile(file, matchAction)  
  return snippets
}

func IterateFile(file string, action func(term string)) error {
  f, err := os.Open(file)
  if err != nil {
    return err
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    for _, currentTerm := range strings.Split(scanner.Text(), " ") {
      action(currentTerm)
    }  
  }
  
  return nil
}

func GetTermMatches(term string) {
  snippets := MatchTermsFromFile("terms.txt", term)
  fmt.Println("Found", len(snippets), "matches for term", term)
  for i, snippet := range snippets {
    fmt.Println("Match", i, ":", snippet)
  }
}

// A more featured example, making use of console I/O, file I/O, and function arguments
func exercise3() {
  fmt.Print("Enter term to search: ")
  term, err := GetSearchTerm()
  if err == nil {
    GetTermMatches(term)
  }
}
