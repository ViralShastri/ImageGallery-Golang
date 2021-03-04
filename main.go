package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
}

func main() {

	http.HandleFunc("/viral", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang As Viral</h1>")
	})

	http.HandleFunc("/", handlerFunc)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
	// }) // Same as Above

	http.ListenAndServe(":3000", nil)
}
