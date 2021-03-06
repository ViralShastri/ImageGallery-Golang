package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(rw, "Welcome!")
	})
	router.GET("/hello/:name", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(rw, "Welcome!, %s", p.ByName("name"))
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/html")
	// 	if r.URL.Path == "/" {
	// 		fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
	// 	} else if r.URL.Path == "/contact-us" {
	// 		fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
	// 	} else {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
	// 	}
	// })

	log.Fatal(http.ListenAndServe(":3000", router))

}
