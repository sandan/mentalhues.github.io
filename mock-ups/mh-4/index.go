package main
import(
  "net/http"
  "html/template"
)

type Story struct{

  Image string //path to static image
  Title string // optional title of the hue
  Description string
  Author string
  Active bool // used for featured stories in carousel.html
}

type Hues struct {
  Featured []Story
  Hue []Story
}
// in-memory for now
var placeholder = "Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit."
var shorter = "Donec id elit non mi porta gravida at eget metus. Nullam id dolor id nibh ultricies vehicula ut id elit."

var data = Hues{
     []Story{
       Story{ Image: "imgs/plants-growing.jpg", Title: "One Step at a time", Description: placeholder, Author: "Anonymous"}, //active
       Story{ Image: "imgs/Nature.jpg", Title: "Reconnecting with Nature", Description: placeholder, Author: "Anonymous"} ,
       Story{ Image: "imgs/space-mountain.jpg", Title: "Exploring what is beyond", Description: placeholder, Author: "Anonymous"},
     },
     []Story{
       Story{ Image: "imgs/ireland-rainbow.jpg", Title: "Welcome to mentalhues!", Description: shorter, Author: "Erin Cole"},
       Story{ Image: "imgs/rainbow-stripe.jpg", Title: "Sensitive and Alone", Description: shorter, Author: "Khirstie Whitaker"},
       Story{ Image: "imgs/floating-candles.jpg", Title: "Survival Mentality", Description: shorter, Author: "Mark Sandan"},
       Story{ Image: "imgs/color-splash-red-blue.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous"},
     },
}

func index(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/mentalhues.html",
    "templates/navbar.html",
    "templates/carousel.html",
    "templates/body.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "index", data)

}
