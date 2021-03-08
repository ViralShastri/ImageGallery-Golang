package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTemplate         *template.Template
	contactTemplate      *template.Template
	faqTemplate          *template.Template
	pageNotFoundTemplate *template.Template
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
	if err := pageNotFoundTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
	if err := homeTemplate.Execute(w, nil); err != nil {
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

	if err := faqTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}

func main() {

	var err error

	homeTemplate, err = template.ParseFiles("views/home.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles("views/contact.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	faqTemplate, err = template.ParseFiles("views/faq.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	pageNotFoundTemplate, err = template.ParseFiles("views/404.gohtml", "views/layouts/footer.gohtml")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)

	// router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFound).GetHandler()
	router.NotFoundHandler = http.HandlerFunc(pageNotFound)

	log.Fatal(http.ListenAndServe(":3000", router))

}
