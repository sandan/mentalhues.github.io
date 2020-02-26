package main
import(
  "net/http"
  "html/template"
)

func code_of_conduct(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/code_of_conduct.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "index", nil)

}
