package handlers

import "net/http"

func GetTestBody(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get url"))
}

func PostSaveSsilka(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from post url"))
}
