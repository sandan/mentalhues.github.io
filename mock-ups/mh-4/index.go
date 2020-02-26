package main
import(
  "net/http"
  "html/template"
  "strings"
  "math/rand"
  "time"
)

type Story struct{

  Image string //path to static image
  Title string // optional title of the hue
  Description string
  Author string
  Body string
  URL string
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
       Story{ Image: "imgs/plants-growing.jpg", Title: "One Step at a time", Description: placeholder, Author: "Anonymous", Body: strings.Repeat(placeholder, 100)}, //active
       Story{ Image: "imgs/Nature.jpg", Title: "Reconnecting with Nature", Description: placeholder, Author: "Anonymous"} ,
       Story{ Image: "imgs/space-mountain.jpg", Title: "Exploring what is beyond", Description: placeholder, Author: "Anonymous"},
     },
     []Story{
       Story{ Image: "imgs/abstract-color-wave-kathy-kurtz.jpg", Title: "Welcome to mentalhues!", Description: shorter, Author: "Erin Cole", Body: strings.Repeat(placeholder + shorter, 50), URL:"/welcome-to-mentalhues"},
       Story{ Image: "imgs/rainbow-stripe.jpg", Title: "Sensitive and Alone", Description: shorter, Author: "Khirstie Whitaker", Body: strings.Repeat(placeholder + shorter, 50), URL: "/sensitive-and-alone"},
       Story{ Image: "imgs/smokies.jpg", Title: "Survival Mentality", Description: shorter, Author: "Mark Sandan", Body: strings.Repeat(placeholder + shorter, 50), URL:"/survival-mentality"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
       Story{ Image: "imgs/mist.jpg", Title: "Hope, Struggle, Patience", Description: shorter, Author: "Anonymous", Body: strings.Repeat(placeholder+shorter, 50), URL:"/hope-struggle-patience"},
     },
}

func getRandomColor(i int) string{

  var letters = []string{"0", "1", "2", "3", "4", "5", "6", "7",
                         "8", "9", "A", "B", "C", "D", "E", "F"}
  color := "#"

  rand.Seed(int64(i) + time.Now().UnixNano())

  for x := 0; x < 6; x++ {
    color += letters[rand.Intn(16)];
  }
  return color
}

func index(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/layout.html",
    "templates/navbar.html",
    "templates/hues_art.html",
    "templates/footer.html",
  }

  funcMap := template.FuncMap{ "rhue": getRandomColor }
  templates := template.New("index").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))
  templates.Execute(writer, data)
}
