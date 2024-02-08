package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.tmpl")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
   path := r.URL.Path
   if strings.HasSuffix(path, "js") {
      w.Header().Set("Content-Type","text/javascript")
   } else {
      w.Header().Set("Content-Type","text/css")
   }
   data, err := os.ReadFile(path[1:])
   if err != nil {
        log.Println(err)
    }
   _, err = w.Write(data)
   if err != nil {
        log.Println(err)
   }
}


