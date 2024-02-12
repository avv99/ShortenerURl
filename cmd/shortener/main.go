package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fadsfa"))
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fadsfa"))
	})

	log.Println("Я завелся")
	http.ListenAndServe(":8080", r)

}
