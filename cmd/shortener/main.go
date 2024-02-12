package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"shortenerurl/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostSaveSsilka(w, r)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTestBody(w, r)
	})

	log.Println("Я завелся")
	http.ListenAndServe(":8080", r)

}
