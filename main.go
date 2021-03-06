package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
		} else if r.URL.Path == "/contact-us" {
			fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
		}
	})
	http.ListenAndServe(":3000", nil)
}
