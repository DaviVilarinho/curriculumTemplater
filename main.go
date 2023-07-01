package main

import (
	"html/template"
	"log"
	"net/http"
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
  Logo: "https://upload.wikimedia.org/wikipedia/en/f/f7/RickRoll.png",
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))
func indexHandler(res http.ResponseWriter, req *http.Request) {
  if err := indexTemplate.Execute(res, index); err != nil {
    log.Fatal("Renderizing error");
  }

}

func curriculumHandler(res http.ResponseWriter, req *http.Request) {
}
