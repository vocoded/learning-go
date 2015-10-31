package main

import (
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
)

func ReadUrl(url string) []byte {
  response, _ := http.Get(url)
  defer response.Body.Close()
  body, _ := ioutil.ReadAll(response.Body)
  
  return body
}

// Download a given file and record how long it takes
func exercise7() {
  var url string
  
  fmt.Print("Enter URL to download: ")
  fmt.Scanln(&url)

  start := time.Now()
  body := ReadUrl(url)
  elapsed := time.Since(start)
  
  fmt.Println("Took", elapsed, "to retrieve", len(body), "bytes")
}
