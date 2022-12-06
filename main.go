package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello! Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `To get in touch, please send an email to <a
	href=\"mailto:noahzajicek@gmail.com\">noahzajicek@gmail.com </a>.`)
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<h1>Frequently asked questions:</h1><p>Here is a list of questions 
	our users commonly ask.</p> `)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 error: page not found :(</h1>")
}

func main() {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(notFound)
	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)
	log.Fatal(http.ListenAndServe(":3000", r))
}
