package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func pageNotFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)

	log.Fatal(http.ListenAndServe(":3000", router))

}
