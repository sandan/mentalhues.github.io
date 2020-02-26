package main
import(
  "net/http"
  "html/template"
)

func guidelines(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/guidelines.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "index", nil)

}
