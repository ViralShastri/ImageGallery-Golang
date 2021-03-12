package main

import (
	"log"
	"net/http"

	"github.com/ViralShastri/ImageGallery-Golang/controllers"
	"github.com/ViralShastri/ImageGallery-Golang/views"
	"github.com/gorilla/mux"
)

var (
	pageNotFoundView *views.View
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(pageNotFoundView.Render(w, nil))
}

func main() {
	statiController := controllers.NewStatic()
	pageNotFoundView = views.NewView("bootstrap", "404")

	usersController := controllers.NewUsers()

	router := mux.NewRouter()

	router.Handle("/", statiController.Home).Methods("GET")
	router.Handle("/contact", statiController.Contact).Methods("GET")
	router.HandleFunc("/signup", usersController.New).Methods("GET")
	router.HandleFunc("/signup", usersController.Create).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(pageNotFound)

	log.Fatal(http.ListenAndServe(":3000", router))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
