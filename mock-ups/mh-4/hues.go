package main
import(
  "net/http"
  "html/template"
)

func GetHueByURL(url string) *Story{
// for now, we just use in-memory hues
// these have attributes for URL
  for _, s := range data.Hue {
    if ("/hue" + s.URL) == url {
      return &s
    }
  }
  return nil
}

func hues(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/story.html",
    "templates/footer.html",
  }

  // get the hue to read
  vals := request.URL.RequestURI()
  hue := GetHueByURL(vals)

  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "index", hue)
}
