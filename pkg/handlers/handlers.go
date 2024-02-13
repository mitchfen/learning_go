package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mitchfen/my_go_site/pkg/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl")
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
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
