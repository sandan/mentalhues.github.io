package main

import (
  "net/http"
  "html/template"
)

func signin(writer http.ResponseWriter, request *http.Request){

  templates := template.Must(template.ParseFiles("templates/signin.html"))
  templates.ExecuteTemplate(writer, "signin", nil)

}
