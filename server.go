package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
  "os"
)

func Home(w http.ResponseWriter, req *http.Request) {
	render(w, "index.html")
}

func About(w http.ResponseWriter, req *http.Request) {
	render(w, "about.html")
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
