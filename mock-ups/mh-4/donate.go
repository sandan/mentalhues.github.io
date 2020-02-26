package main

import (
  "net/http"
  "html/template"
)

func donate(writer http.ResponseWriter, request *http.Request){

  templates := template.Must(template.ParseFiles("templates/donate.html"))
  templates.ExecuteTemplate(writer, "donate", nil)

}
