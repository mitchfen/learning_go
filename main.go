package main

import (
	"fmt"
	"net/http"
)

const PortNumber string = "5000"

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/about", About)
	http.HandleFunc("/static/", staticHandler)
	err := http.ListenAndServe(":"+PortNumber, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}
