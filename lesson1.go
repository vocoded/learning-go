package main

import (
  "fmt"
  "os" 
  "bufio"
  "bytes"
)

func GetSearchTerm() string {
  var reader = bufio.NewReader(os.Stdin) 
  var term, _ = reader.ReadString('\n')
  return term
}

func ScanSearchTerm() string {
  var term string
  fmt.Scanln(&term)
  return term
}

func MatchTerms(term string) (int, error) {
  var count = 0;
  f, err := os.Open("terms.txt")
  if err != nil {
    return 0, err
  }
  
  defer f.Close()
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    if bytes.Index(scanner.Bytes(), []byte(term)) >= 0 {
      count = count + 1
    }
  }
  
  return count, nil
}

func main() {
  fmt.Print("Enter term to search: ")
  var term = ScanSearchTerm()
  var count, _ = MatchTerms(term)
  fmt.Println("Found", count, "matches for term", term)
}
