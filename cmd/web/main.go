package main

import (
	"fmt"
	"net/http"

	"github.com/mitchfen/my_go_site/pkg/handlers"
)

const PortNumber string = "5000"

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/static/", handlers.StaticHandler)
	err := http.ListenAndServe(":"+PortNumber, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}
