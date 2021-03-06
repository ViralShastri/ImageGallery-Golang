package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "Welcome!")
}

func faq(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html")

	fmt.Fprint(rw, `<ul>
	<li>FAQ 1</li>
	<li>FAQ 2</li>
	<li>FAQ 2</li>
	</ul>
	`)
}

func contact(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html")
	fmt.Fprint(rw, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
}

func main() {

	router := httprouter.New()
	router.GET("/", home)
	router.GET("/faq", faq)
	router.GET("/contact", contact)
	router.NotFound = http.HandlerFunc(pageNotFound)
	http.ListenAndServe(":3000", router)

}
