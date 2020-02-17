package main

import(
  "os"
  "fmt"
  "encoding/json"
)

type Configuration struct {
    Address      string
    ReadTimeout  int64
    WriteTimeout int64
    Static       string
}

var config Configuration

func init(){
    file, err := os.Open("config.json")
    if err != nil {
        fmt.Println("Cannot open config file", err)
    }
    decoder := json.NewDecoder(file)
    config = Configuration{}
    err = decoder.Decode(&config)
    if err != nil {
        fmt.Println("Cannot get configuration from file", err)
    }

}
