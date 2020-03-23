package main

import(
  "net/http"
  "log"
  "time"
  "os"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
  "encoding/json"
)

type Configuration struct {
    Address      string
    ReadTimeout  int64
    WriteTimeout int64
    Static       string
}
var config Configuration
var Db *sql.DB

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

    Db, err = sql.Open("postgres", "dbname=chatchat sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    return
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
  mux.HandleFunc("/donate", donate)

  mux.HandleFunc("/signin", signin)
  mux.HandleFunc("/authn", auth)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", account)
  mux.HandleFunc("/signout", signout)

  //mux.HandleFunc("/share", share)
  //mux.HandleFunc("/contributors", contrib)
  //mux.HandleFunc("/podcast", podcast)
  //mux.HandleFunc("/publishing", volumes)
  //mux.HandleFunc("/web-series", web-series)
  //mux.HandleFunc("/err", err)


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
