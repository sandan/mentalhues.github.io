package main

import (
  "net/http"
  "time"
)


func main(){

  mux := http.NewServeMux()
  files := http.FileServer(http.Dir(config.Static))
  mux.Handle("/static/", http.StripPrefix("/static/", files))

  mux.HandleFunc("/", index)
  mux.HandleFunc("/signin", signin)
  mux.HandleFunc("/about", about)
  //mux.HandleFunc("/signup", signup)

  //http.HandleFunc("/err", err)
  //http.HandleFunc("/login", login)
  //http.HandleFunc("/logout", logout)
  //http.HandleFunc("/authenticate", logout)

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
