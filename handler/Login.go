package handler

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "" || password == "" {
			ShowErrorPage(w, "Username and password are required", http.StatusBadRequest)
			return
		}
	}
}
