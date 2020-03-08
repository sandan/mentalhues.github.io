package main

import(
  "html/template"
  "net/http"
)

type Info struct{
  Banner Banner
  Hues Hues
}

type Banner struct{
  Content []string
  Active string
  Display string
  Body string
}

var(

 banner = Banner{
  Content: []string{
    "Community",
    "Hues",
    "Gallery",
    "Blog",
    "Web Series",
    "Podcast",
    "Store",
  },
}

 hues_info = Hues{
  Featured: []Hue{
      Hue{
        Images: "imgs/wall-1.jpg",
        Title: "Cold",
        Body: "This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
        Id: 1,
      },
      Hue{
        Images: "imgs/aurora-borealis.jpg",
        Title: "Cold",
        Body: "This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
        Id: 2,
      },
      Hue{
        Images: "imgs/color-splash-red-blue.jpg",
        Title: "Cold",
        Body: "This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
        Id: 3,
      },
  },
}

 info = Info{
    Banner: banner,
    Hues: hues_info,
  }
)

func index(writer http.ResponseWriter, request *http.Request){
  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/banner.html",
    "templates/users-featured.html",
    "templates/wall.html",
    "templates/ad.html",
    "templates/footer.html",
  }

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Community"
  info.Banner.Body = "Feel their Stories"
  info.Banner.Display = "/static/imgs/ocean-dusk.jpg"
  templates.Execute(writer, info)
}

func gallery(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/banner.html",
    "templates/users-featured.html",
    "templates/gallery.html",
    "templates/ad.html",
    "templates/footer.html",
  }

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Gallery"
  info.Banner.Body = "See their Stories"
  info.Banner.Display = "/static/imgs/rainbow-stripe.jpg"
  templates.Execute(writer, info)
}

func hues(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/banner.html",
    "templates/users-featured.html",
    "templates/hues-featured.html",
    "templates/hues.html",
    "templates/ad.html",
    "templates/footer.html",
  }

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Hues"
  info.Banner.Body = "Read their Stories"
  info.Banner.Display = "/static/imgs/mist.jpg"
  templates.Execute(writer, info)
}

func signin(writer http.ResponseWriter, request *http.Request){}
func signup(writer http.ResponseWriter, request *http.Request){}
func signout(writer http.ResponseWriter, request *http.Request){}
func donate(writer http.ResponseWriter, request *http.Request){}
func share(writer http.ResponseWriter, request *http.Request){}
func wall(writer http.ResponseWriter, request *http.Request){}
func authenticate(writer http.ResponseWriter, request *http.Request){}

// static
func about(writer http.ResponseWriter, request *http.Request){}
func guidelines(writer http.ResponseWriter, request *http.Request){}
func conduct(writer http.ResponseWriter, request *http.Request){}
func err404(writer http.ResponseWriter, request *http.Request){}
