package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
  //"github.com/golang-jwt/jwt/v5"
	// "golang.org/x/exp/slices"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/curriculum", curriculumHandler)
  http.HandleFunc("/login", loginHandler)
  //http.HandleFunc("/registration", registrationHandler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type IndexPage struct {
  Title string
  Logo string
}

var index = &IndexPage{
  Title: "Welcome to Curriculum Templater",
  Logo: "://upload.wikimedia.org/wikipedia/en/f/f7/RickRoll.png",
}

type User struct {
  UserName string
  Password string
}

var users map[string]User = make(map[string]User)
func loginHandler(res http.ResponseWriter, req *http.Request) {
  parseErr := req.ParseForm()
  usernameFromForm, ok := req.Form["username"]
  if !ok || parseErr != nil {
    log.Println(parseErr)
    log.Println(ok)
    res.WriteHeader(http.StatusBadRequest)
    res.Write([]byte("Bad Request Login"))
  }
  username, ok := users[strings.Join(usernameFromForm, "")] 
  if ok && strings.Join(req.Form["password"], "") == username.Password {
      res.WriteHeader(http.StatusOK)
      res.Write([]byte("Succesfully logged in"))
      return
  }
  res.WriteHeader(http.StatusUnauthorized)
  res.Write([]byte("Login did not work"))
  return
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))
func indexHandler(res http.ResponseWriter, req *http.Request) {
  if err := indexTemplate.Execute(res, index); err != nil {
    log.Println("Renderizing error");
  }
}

type Field struct {
  Name string
  Content string
}

type CurriculumPage struct {
  Name string
  Fields []Field
}

var curriculumTemplate = template.Must(template.ParseFiles("curriculum.tmpl"))
func curriculumHandler(res http.ResponseWriter, req *http.Request) {
  if err := req.ParseForm(); err != nil {
    log.Println("Could not parse form")
    return
  }
  fields := []Field{}
  for key, val := range req.PostForm {
    field := Field{Name: key, Content: strings.Join(val, "")}
    fields = append(fields, field)
  }
  dataParsed := &CurriculumPage{
    Name: strings.Join(req.Form["name"], ""),
    Fields: fields,
  }
  if err := curriculumTemplate.Execute(res, dataParsed); err != nil {
    log.Println("Parsing Error")
  }
}
