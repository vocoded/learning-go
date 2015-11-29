package main

import (
  "fmt"
  "sort"
)

// Apply a sort to the unique terms found in the given file
func exercise6() {
  counter := FileTermFinder {"terms.txt", make(map[string]bool)}  
  words := counter.Find()
  sort.Strings(words)
  for _, w := range words {
    fmt.Println(w)
  }
}
