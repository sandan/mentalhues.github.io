package main
import(
  "net/http"
  "html/template"
)

func about(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/about.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "index", nil)

}
