package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Я завелся")
	http.ListenAndServe(":8080", http.DefaultServeMux)

}
