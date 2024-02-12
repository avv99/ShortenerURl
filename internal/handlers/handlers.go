package handlers

import "net/http"

func GetUrlByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get url"))
}

func PostSaveUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from post url"))
}
