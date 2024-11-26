package handler

import (
	"net/http"

	database "Forum/dataBase"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		if email == "" || username == "" || password == "" {
			ShowErrorPage(w, "All fields are required", http.StatusBadRequest)
			return
		}
		// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		// if err != nil {
		// 	ShowErrorPage(w, "Error hashing password", http.StatusInternalServerError)
		// 	return
		// }

		_, err := database.Db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", email, username, password)
		if err != nil {
			ShowErrorPage(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	ShowErrorPage(w, "Invalid request method", http.StatusMethodNotAllowed)
}
