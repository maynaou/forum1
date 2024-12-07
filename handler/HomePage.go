package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ShowErrorPage(w, "Page not found", http.StatusNotFound)
		return
	}

	cookies := r.Cookies()
	if len(cookies) != 0 {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		CurrentUser := cookie.Value
		fmt.Println(CurrentUser)

		http.Redirect(w, r, "/forum", http.StatusSeeOther)
		return
	}

	err := templates.ExecuteTemplate(w, "homePage.html", nil)
	if err != nil {
		ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
