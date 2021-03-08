package main

import (
	"log"
	"net/http"

	"github.com/ViralShastri/usegolang/views"
	"github.com/gorilla/mux"
)

var (
	homeView         *views.View
	contactView      *views.View
	faqView          *views.View
	pageNotFoundView *views.View
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
	if err := pageNotFoundView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, `<ul>
	// <li>FAQ 1</li>
	// <li>FAQ 2</li>
	// <li>FAQ 2</li>
	// </ul>
	// `)

	if err := faqView.Template.Execute(w, nil); err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}

}

func main() {

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	faqView = views.NewView("views/faq.gohtml")
	pageNotFoundView = views.NewView("views/404.gohtml")

	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)

	// router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFound).GetHandler()
	router.NotFoundHandler = http.HandlerFunc(pageNotFound)

	log.Fatal(http.ListenAndServe(":3000", router))

}
