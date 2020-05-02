package main

import (
  "fmt"
  "os"
  "encoding/json"
)

type capability struct {
  Code string  
  Language string
}

func main(){

  json_blob := os.Args[1]
  fmt.Println(json_blob)
  var caps capability

  if err := json.Unmarshal([]byte(json_blob),&caps) ; err != nil {
    fmt.Println(err)
  }
  fmt.Println(caps)
}
