package main

import (
	"log"
	"net/http"

	"github.com/ViralShastri/usegolang/controllers"
	"github.com/ViralShastri/usegolang/views"
	"github.com/gorilla/mux"
)

var (
	homeView         *views.View
	contactView      *views.View
	faqView          *views.View
	pageNotFoundView *views.View
	// signUpView       *views.View
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>404 Page Not Found</h1><p>We couldn't find the page you're looking for :(</p>")
	// if err := pageNotFoundView.Template.Execute(w, nil); err != nil {
	// 	panic(err)
	// }

	// if err := pageNotFoundView.Template.ExecuteTemplate(w, pageNotFoundView.Layout, nil); err != nil {
	// 	panic(err)
	// }

	must(pageNotFoundView.Render(w, nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Welcome to my First Web Application in GoLang</h1>")
	// if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
	// 	panic(err)
	// }

	must(homeView.Render(w, nil))

}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, `<ul>
	// <li>FAQ 1</li>
	// <li>FAQ 2</li>
	// <li>FAQ 2</li>
	// </ul>
	// `)

	// if err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil); err != nil {
	// 	panic(err)
	// }

	must(faqView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "To get in touch, please send us a email at <a href=\"mailto:support@github.com\">support@github.com</a>.")
	// if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
	// 	panic(err)
	// }
	must(contactView.Render(w, nil))

}

// func signUp(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	must(signUpView.Render(w, nil))
// }

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	pageNotFoundView = views.NewView("bootstrap", "views/404.gohtml")
	// signUpView = views.NewView("bootstrap", "views/signup.gohtml")

	usersController := controllers.NewUsers()

	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/faq", faq)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", usersController.New)

	// router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFound).GetHandler()
	router.NotFoundHandler = http.HandlerFunc(pageNotFound)

	log.Fatal(http.ListenAndServe(":3000", router))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
