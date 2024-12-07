package handler

import (
	"net/http"

	database "Forum/dataBase"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		ShowErrorPage(w, "Page Not found", http.StatusNotFound)
		return
	}
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

		var user_id int
		err := database.Db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ? ", username, password).Scan(&user_id)
		if err != nil {
			ShowErrorPage()
		}

		http.Redirect(w, r, "/forum", http.StatusFound)
	}

	ShowErrorPage(w,"Methode not Allowed",htt)
}
