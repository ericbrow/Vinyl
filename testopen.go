package main

import (
  "fmt"
//  "io/ioutil"
  "os"
)

func main() {

  xmlFile, err := os.Open("short_artists.xml")
  if err != nil {
      fmt.Println(err)
  }

  fmt.Println("Successfully opened artists file")
  defer xmlFile.Close()
}
