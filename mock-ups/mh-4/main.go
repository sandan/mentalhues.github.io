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
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/about", about)
  mux.HandleFunc("/donate", donate)
  mux.HandleFunc("/code_of_conduct", code_of_conduct)
  mux.HandleFunc("/guidelines", guidelines)
  mux.HandleFunc("/hue/", hues) // handle subpath
  //mux.HandleFunc("/contributors", contrib)
  //mux.HandleFunc("/podcast", gallery)
  //mux.HandleFunc("/publishing", volumes)
  //mux.HandleFunc("/web-series", web_series)
  //http.HandleFunc("/err", err)
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
