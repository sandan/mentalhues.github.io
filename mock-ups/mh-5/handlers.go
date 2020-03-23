package main

import(
  "html/template"
  "log"
  "net/http"
)

type Info struct{
  Banner Banner
  Hues Hues
  Session bool
  Wall []Thread
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
        Images: []string{"imgs/backtoback.jpg"},
        Title: "Cold",
        Body: "This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
        Id: 1,
      },
      Hue{
        Images: []string{"imgs/people-talk-banner.jpg"},
        Title: "Cold",
        Body: "This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.",
        Id: 2,
      },
      Hue{
        Images: []string{"imgs/watercolor.jpg"},
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

  feed, err := Threads() //get all wall threads
  if err != nil{
    log.Println(err)
  }
  files := []string{
    "templates/index.html",
  }

  // check session
  _, err = session(writer, request)
  if err != nil{ // then user did not authenticate/set cookie
    files = append(files, "templates/public.navbar.html")
    info.Session = false
  } else {
    files = append(files, "templates/private.navbar.html")
    info.Session = true
  }

  files = append(files,
    "templates/body.html",
    "templates/banner.html",
    "templates/users-featured.html",
    "templates/wall-layout.html",
    "templates/wall.html",
    "templates/ad.html",
    "templates/footer.html",
  )

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Community"
  info.Banner.Body = "Feel their Stories"
  if info.Session{
    info.Banner.Body = "Feel their Stories, signed in user"
  }
  info.Banner.Display = "/static/imgs/color-splash-red-blue.jpg"
  info.Wall = feed
  templates.Execute(writer, info)
}

func gallery(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
  }

  // check session
  _, err := session(writer, request)
  if err != nil{ // then user did not authenticate/set cookie
    files = append(files, "templates/public.navbar.html")
    info.Session = false
  } else {
    files = append(files, "templates/private.navbar.html")
    info.Session = true
  }

  files = append(files,
    "templates/body.html",
    "templates/banner.html",
    "templates/gallery.html",
    "templates/footer.html",
  )

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Gallery"
  info.Banner.Body = "See their Stories"
  info.Banner.Display = "/static/imgs/color-splash-red-blue.jpg"
  templates.Execute(writer, info)
}

func hues(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
  }

  // check session
  _, err := session(writer, request)
  if err != nil{ // then user did not authenticate/set cookie
    files = append(files, "templates/public.navbar.html")
    info.Session = false
  } else {
    files = append(files, "templates/private.navbar.html")
    info.Session = true
  }

  files = append(files,
    "templates/body.html",
    "templates/banner.html",
    "templates/hues-featured.html",
    "templates/hues.html",
    "templates/footer.html",
  )

  funcMap := template.FuncMap{ "rhue": getRandomColor, "lower": lower }
  templates := template.New("layout").Funcs(funcMap)
  templates = template.Must(templates.ParseFiles(files...))

  info.Banner.Active = "Hues"
  info.Banner.Body = "Read their Stories"
  info.Banner.Display = "/static/imgs/color-splash-red-blue.jpg"
  templates.Execute(writer, info)
}

func share(writer http.ResponseWriter, request *http.Request){}

/* auth & identity handlers */
// The signin.html has a form that redirects to /auth using POST
func signin(writer http.ResponseWriter, request *http.Request){
  templates := template.Must(template.ParseFiles("templates/signin.html"))
  templates.ExecuteTemplate(writer, "signin", nil)
}

// sign up page
// redirects to account handler function via POST
func signup(writer http.ResponseWriter, request *http.Request){
  templates := template.Must(template.ParseFiles("templates/signup.html"))
  templates.ExecuteTemplate(writer, "signup", nil)
}

// create account
func account(writer http.ResponseWriter, request *http.Request){
    err := request.ParseForm()
    if err != nil {
        log.Println(err, "Cannot parse form")
    }
    user := User{
        Name:     request.PostFormValue("name"),
        Email:    request.PostFormValue("email"),
        Password: request.PostFormValue("password"),
    }
    //TODO: check for avatar if image field is set; save it on disk and place the path to it in User.Image
    if err := user.Create(); err != nil {
        log.Println(err, "Cannot create user")
    }
    http.Redirect(writer, request, "/signin", 302)
}

// User can only sign out if they have signed in already
// Once a user signs in, there should be a cookie with the
// uuid of their session kept on their browser.
func signout(writer http.ResponseWriter, request *http.Request){
    cookie, err := request.Cookie("my_cookie")
    if err != http.ErrNoCookie {
        log.Println(err, "Failed to get cookie")
        session := Session{Uuid: cookie.Value}
        session.DeleteByUUID()
    }
    http.Redirect(writer, request, "/", 302)
}

func donate(writer http.ResponseWriter, request *http.Request){
  templates := template.Must(template.ParseFiles("templates/donate.html"))
  templates.ExecuteTemplate(writer, "donate", nil)
}

//   User authenticates during signin 
//   Application calls the session function to check for the cookie
func auth(writer http.ResponseWriter, request *http.Request){
    err := request.ParseForm()
    user, err := UserByEmail(request.PostFormValue("email"))
    if err != nil{
    }
    if user.Password == Encrypt(request.PostFormValue("password")) {
        session, err := user.CreateSession()
        if err != nil {
            log.Println(err, "Cannot create session")
        }
        // TODO: check use-cookie is set from the "Remember me" in form
        // https://computer.howstuffworks.com/cookie.htm
        // this gets stored on the client's disk
        // subsequent requests from the client send any cookies we ask for by name, for only cookies we've sent in Response headers
        // https://security.stackexchange.com/questions/49636/can-a-webpage-read-another-pages-cookies
        cookie := http.Cookie{
            Name:     "my_cookie",
            Value:    session.Uuid,
            HttpOnly: true,
        }
        http.SetCookie(writer, &cookie)
        http.Redirect(writer, request, "/", 302) // the session function will check for the cookie that was set
    } else {
        http.Redirect(writer, request, "/signin", 302)
    }
}

// static
func about(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/about.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "layout", nil)
}
func guidelines(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/guidelines.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "layout", nil)
}
func conduct(writer http.ResponseWriter, request *http.Request){

  files := []string{
    "templates/index.html",
    "templates/navbar1.html",
    "templates/conduct.html",
    "templates/footer.html",
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(writer, "layout", nil)
}
func err404(writer http.ResponseWriter, request *http.Request){

}
