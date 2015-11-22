package main

import (
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
)

func ReadUrl(url string) []byte {
  response, err := http.Get(url)
  if err == nil {
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
  
    return body
  }

  return make([]byte, 0, 0)
}

func TimeReadUrl(url string) {
  start := time.Now()
  body := ReadUrl(url)
  elapsed := time.Since(start)
  
  fmt.Println("Took", elapsed, "to retrieve", len(body), "bytes")  
}

// Download a given file and record how long it takes
func exercise7() {
  var url string
  
  fmt.Print("Enter URL to download: ")
  if _, err := fmt.Scanln(&url); err == nil {
    TimeReadUrl(url)
  }
}
