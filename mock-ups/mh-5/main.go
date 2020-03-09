package main

import(
  "net/http"
  "time"
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

func main(){
  mux := http.NewServeMux()
  files := http.FileServer(http.Dir(config.Static))
  mux.Handle("/static/", http.StripPrefix("/static/", files))

  mux.HandleFunc("/", index)
  mux.HandleFunc("/community", index)
  mux.HandleFunc("/gallery", gallery)
  mux.HandleFunc("/hues", hues)

  mux.HandleFunc("/about", about)
  mux.HandleFunc("/guidelines", guidelines)
  mux.HandleFunc("/conduct", conduct)

  mux.HandleFunc("/signin", signin)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/donate", donate)

  //mux.HandleFunc("/share", share)
  //mux.HandleFunc("/contributors", contrib)
  //mux.HandleFunc("/podcast", podcast)
  //mux.HandleFunc("/publishing", volumes)
  //mux.HandleFunc("/web-series", web-series)
  //http.HandleFunc("/err", err)
  //http.HandleFunc("/logout", logout)
  //http.HandleFunc("/authenticate", auth)

  // starting up the server
  server := &http.Server{
      Addr:           config.Address,
      Handler:        mux,
      ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
      WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
      MaxHeaderBytes: 1 << 20,
  }
  server.ListenAndServe()
}
