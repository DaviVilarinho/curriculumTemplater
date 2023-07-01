package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	// "golang.org/x/exp/slices"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/curriculum", curriculumHandler)
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

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))
func indexHandler(res http.ResponseWriter, req *http.Request) {
  if err := indexTemplate.Execute(res, index); err != nil {
    log.Fatal("Renderizing error");
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
  req.ParseForm()
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
    log.Fatal("Parsing Error")
  }
}
