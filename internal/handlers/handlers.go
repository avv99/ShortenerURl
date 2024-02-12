package handlers

import "net/http"

func GetURLByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get url"))
}

func PostSaveURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from post url"))
}
