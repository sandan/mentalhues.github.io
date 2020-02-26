package main

import (
  "net/http"
  "html/template"
)

func signup(writer http.ResponseWriter, request *http.Request){

  templates := template.Must(template.ParseFiles("templates/signup.html"))
  templates.ExecuteTemplate(writer, "signup", nil)

}
