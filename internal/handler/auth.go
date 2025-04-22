package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Print(username, password)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Received username: %s have %s pass", username, password)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {

}
