package main

import (
  "fmt"
  "os"
  "encoding/json"
  "github.com/itsjwala/locus/runner/languages"
)

type capability struct {
  code     string   `json:"code"`
  language string   `json:"language"`
}

// type result struct {
//   Stdout string     `json:"stdout"`
//   Stderr string     `json:"stderr"`
//   Error  string     `json:"error"`
// }
func main(){


  caps := &capability{}

  if err := json.NewDecoder(os.Stdin).Decode(caps) ; err != nil {
     fmt.Println(err)
     panic(err)
  }

  fmt.Println(caps)
  fmt.Println("Runnning")


  if ! languages.IsSupported(caps.language) {
    panic("Language is not supported")
  }


  languages.Run(caps.language,caps.code)

}
