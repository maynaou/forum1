package handler

import (
	"net/http"

	database "Forum/dataBase"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		ShowErrorPage(w, "Page Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err := templates.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			ShowErrorPage(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		var userID int
		err := database.Db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&userID)
		if err != nil {
			ShowErrorPage(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/forum", http.StatusFound)
		return

	default:
		ShowErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
